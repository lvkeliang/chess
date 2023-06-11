package api

import (
	"Gone/model"
	"Gone/service"
	"Gone/uitl"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func CreateBoard(h *model.Hub, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取用户ID和房间名
		//从token获取uID,token已经保证了用户必须登录
		tempID, exists := c.Get("ID")

		//处理uID

		if !exists {
			util.RespDidNotLogin(c)
			return
		}

		ID, ok := tempID.(uint)
		if !ok {
			fmt.Println(ok)
			util.RespUnexceptedError(c)
			return
		}

		BoardName := c.PostForm("board_name")
		newBoard, err := service.Reset()
		if err != nil {
			c.JSON(404, gin.H{
				"message": "board create failed",
			})
			return
		}

		// 在数据库中创建一个新的房间记录
		board := model.Board{Name: BoardName, WhiteID: ID,
			State: newBoard,
			Turn:  true,
		}
		if db.Create(&board).Error != nil {
			c.JSON(404, gin.H{
				"message": "room create failed",
			})
			return
		}
		// 返回成功的响应
		c.JSON(200, gin.H{
			"message":  "create board success",
			"board_id": board.ID,
		})
	}
}

func JoinBoard(h *model.Hub, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取用户ID和房间名
		//从token获取uID,token已经保证了用户必须登录
		tempID, exists := c.Get("ID")
		//处理uID

		if !exists {
			util.RespDidNotLogin(c)
			return
		}

		ID, ok := tempID.(uint)
		if !ok {
			fmt.Println(ok)
			util.RespUnexceptedError(c)
			return
		}

		boardID := c.PostForm("boardid")

		// 在数据库中查询房间记录
		var board model.Board
		var audience map[uint]bool
		if db.First(&board, boardID).Error == nil {

			if board.WhiteID == ID || board.BlackID == ID { // 重复加入
				c.JSON(200, gin.H{
					"message": "join room repeated",
				})
				return
			}

			if board.WhiteID == 0 {
				board.WhiteID = ID
			} else if board.BlackID == 0 {
				board.BlackID = ID
			} else {
				err := json.Unmarshal([]byte(board.AudienceID), &audience)
				if err != nil {
					c.JSON(404, gin.H{
						"message": "join room as audience failed",
					})
					log.Println(err)
					return
				}
				audience[ID] = true
				tempAudienceID, err := json.Marshal(audience)
				if err != nil {
					c.JSON(404, gin.H{
						"message": "join room failed",
					})
					log.Println(err)
					return
				}
				board.AudienceID = string(tempAudienceID)
			}

			result := db.Save(&board)
			if result.Error != nil {
				log.Println(result.Error)
			}

			// 返回成功的响应
			c.JSON(200, gin.H{
				"message":    "join board success",
				"board_name": board.Name,
			})
		} else {
			c.JSON(404, gin.H{
				"message": "room not found",
			})
		}
	}
}

// upgrader是一个websocket升级器，它检查连接的来源
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 为了简单起见，允许所有来源
	},
}

func BoardWebSocket(h *model.Hub, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//从token获取uID,token已经保证了用户必须登录
		tempID, exists := c.Get("ID")
		//处理uID

		if !exists {
			util.RespDidNotLogin(c)
			return
		}
		ID, ok := tempID.(uint)
		if !ok {
			fmt.Println(ok)
			util.RespUnexceptedError(c)
			return
		}

		boardIDtemp := c.Param("boardid")
		//处理输入参数并校验合法性
		if len(boardIDtemp) < 1 {
			util.RespAIDError(c)
			return
		}

		boardIDint, err := strconv.Atoi(boardIDtemp)
		if err != nil {
			fmt.Println("boardID转换出错: ", err.Error())
			util.RespUnexceptedError(c)
			return
		}

		playingFlag := 0

		boardID := uint(boardIDint)
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil) // 将请求升级为websocket连接
		if err != nil {
			log.Println(err)
			return
		}
		h.Register <- map[uint]*websocket.Conn{boardID: conn} // 将新连接注册到hub

		for {
			messageType, message, err := conn.ReadMessage() // 从连接中读取一条消息

			if err != nil {
				log.Println(err)
				//log.Println(message)
				h.Unregister <- map[uint]*websocket.Conn{boardID: conn} // 如果有错误，就将连接从hub中注销
				break
			}

			if messageType == websocket.TextMessage { // 如果消息类型是文本
				log.Printf("Received: %s\n", message) // 将收到的消息打印到控制台

				var ready model.Ready                 // 创建一个ready实例来存储消息
				err = json.Unmarshal(message, &ready) // 将消息解析到ready实例中
				if err != nil {
					log.Println(err)
					continue
				}

				ready.UserID = ID
				ready.BoardID = boardID

				var board model.Board                                // 创建一个board实例来更新状态和轮次
				result := db.First(&board, ready.BoardID)            // 根据ready的board ID查找board
				if result.Error != nil || result.RowsAffected == 0 { // 如果没有找到或者有错误
					log.Println(result.Error)
					continue
				}

				if board.WhiteID == ready.UserID {
					board.WhiteReady = ready.Ready
				} else if board.BlackID == ready.UserID {
					board.BlackReady = ready.Ready
				}

				if board.WhiteReady == true && board.BlackReady == true {
					board.Playing = true
				}

				result = db.Save(&board) // 将棋盘保存到数据库中
				if result.Error != nil {
					log.Println(result.Error)
					continue
				}

				h.UpdateBoard <- board

				if board.Playing == true {
					playingFlag = 1
					break
				}
			}
		}

		for playingFlag == 1 {
			messageType, message, err := conn.ReadMessage() // 从连接中读取一条消息

			if err != nil {
				log.Println(err)
				//log.Println(message)
				h.Unregister <- map[uint]*websocket.Conn{boardID: conn} // 如果有错误，就将连接从hub中注销
				break
			}

			if messageType == websocket.TextMessage { // 如果消息类型是文本
				log.Printf("Received: %s\n", message) // 将收到的消息打印到控制台

				var move model.Move                  // 创建一个move实例来存储消息
				err = json.Unmarshal(message, &move) // 将消息解析到move实例中
				if err != nil {
					log.Println(err)
					continue
				}

				move.UserID = ID
				move.BoardID = boardID

				result := db.Create(&move) // 将move保存到数据库中
				if result.Error != nil {
					log.Println(result.Error)
					continue
				}

				var board model.Board                                // 创建一个board实例来更新状态和轮次
				result = db.First(&board, move.BoardID)              // 根据move的board ID查找board
				if result.Error != nil || result.RowsAffected == 0 { // 如果没有找到或者有错误
					log.Println(result.Error)
					continue
				}

				isValid, event, err := service.IsValidMove(move, board)
				if err != nil {
					log.Println(err)
					continue
				}

				if isValid {
					if event == util.WhiteWin {
						board.Winner = 1
						playingFlag = 0
					} else if event == util.BlackWin {
						board.Winner = 0
						playingFlag = 0
					}

					board.State, err = updateState(board.State, move.From, move.To) // 通过应用移动来更新状态
					if err != nil {
						log.Println(err)
						continue
					}
					board.Turn = !board.Turn // 通过翻转它来更新回合
					result = db.Save(&board) // 将棋盘保存到数据库中
					if result.Error != nil {
						log.Println(result.Error)
						continue
					}
				}

				h.UpdateBoard <- board
			}
		}
	}
}

// updateState 通过将一个位置的棋子移动到另一个位置来更新棋盘状态
// 这是一个非常简单的实现，没有检查任何象棋规则或逻辑，你应该根据自己的需要改进它
func updateState(state []byte, from string, to string) ([]byte, error) {
	tempState := make([][]string, 64)
	err := json.Unmarshal(state, &tempState)
	if err != nil {
		return nil, err
	}

	fromIndexX := int(from[0] - 'a') // 计算状态字符串中 from 位置的索引
	fromIndexY := int(from[1] - '1')
	toIndexX := int(to[0] - 'a')
	toIndexY := int(to[1] - '1') // 计算状态字符串中 to 位置的索引
	log.Printf("fx: %v\nfy: %v\ntx: %v\nty: %v\n", fromIndexX, fromIndexY, toIndexX, toIndexY)
	service.PrintBoard(state)

	piece := tempState[fromIndexY][fromIndexX] // 从 from 位置获取棋子，例如 "P"
	tempState[fromIndexY][fromIndexX] = ""     // 在状态字符串中将 from 位置设置为空（""）
	tempState[toIndexY][toIndexX] = piece      // 在状态字符串中将 to 位置设置为棋子

	state, err = json.Marshal(tempState)
	if err != nil {
		return nil, err
	}

	log.Printf("piece: %v\n", piece)
	err = service.PrintBoard(state)
	if err != nil {
		return nil, err
	}

	return state, err // 返回更新后的状态字符串
}
