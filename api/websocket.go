package api

// upgrader是一个websocket升级器，它检查连接的来源
/*var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 为了简单起见，允许所有来源
	},
}*/

// HandleWebSocket处理websocket请求
/*func HandleWebSocket(h *model.Hub, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil) // 将请求升级为websocket连接
		if err != nil {
			log.Println(err)
			return
		}
		h.Register <- conn // 将新连接注册到hub

		for {
			messageType, message, err := conn.ReadMessage() // 从连接中读取一条消息

			if err != nil {
				log.Println(err)
				//log.Println(message)
				h.Unregister <- conn // 如果有错误，就将连接从hub中注销
				break
			}
			if messageType == websocket.TextMessage { // 如果消息类型是文本（你也可以处理二进制消息）
				log.Printf("Received: %s\n", message) // 将收到的消息打印到控制台
				h.Broadcast <- message                // 将消息广播给所有客户端

				var move model.Move                  // 创建一个move实例来存储消息
				err = json.Unmarshal(message, &move) // 将消息解析到move实例中
				if err != nil {
					log.Println(err)
					continue
				}
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
				board.State = updateState(board.State, move.From, move.To) // 通过应用移动来更新状态
				board.Turn = !board.Turn                                   // 通过翻转它来更新回合
				result = db.Save(&board)                                   // 将棋盘保存到数据库中
				if result.Error != nil {
					log.Println(result.Error)
					continue
				}
			}
		}
	}
}
*/
// updateState 通过将一个位置的棋子移动到另一个位置来更新棋盘状态
// 这是一个非常简单的实现，没有检查任何象棋规则或逻辑，你应该根据自己的需要改进它
/*func updateState(state string, from string, to string) string {
	fromIndex := (8-int(from[1]-'0'))*8 + int(from[0]-'a')      // 计算状态字符串中 from 位置的索引，例如 "e2" -> 52
	toIndex := (8-int(to[1]-'0'))*8 + int(to[0]-'a')            // 计算状态字符串中 to 位置的索引，例如 "e4" -> 36
	piece := state[fromIndex]                                   // 从 from 位置获取棋子，例如 "P"
	state = state[:fromIndex] + "." + state[fromIndex+1:]       // 在状态字符串中将 from 位置设置为空（"."）
	state = state[:toIndex] + string(piece) + state[toIndex+1:] // 在状态字符串中将 to 位置设置为棋子
	return state                                                // 返回更新后的状态字符串
}
*/
