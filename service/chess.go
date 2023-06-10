package service

import (
	"Gone/model"
	"fmt"
	"github.com/goccy/go-json"
	"log"
)

func abs(x int) int { // 求绝对值的辅助函数
	if x < 0 {
		return -x
	}
	return x
}

var WhitePieses = []string{"PAWN", "PAWN", "PAWN", "PAWN", "PAWN", "PAWN", "PAWN", "PAWN", "ROOK"}
var BlackPieces = []string{"pawn", "pawn", "pawn", "pawn", "pawn", "pawn", "pawn", "pawn", "rook"}

// 判断字符串切片s中是否包含某个字符串x
func InSlice(s []string, x string) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func Reset() (newBoard []byte, err error) {
	board := [8][8]string{

		{"ROOK", "KNIGHT", "BISHOP", "QUEEN", "KING", "BISHOP", "KNIGHT", "ROOK"},
		{"PAWN", "PAWN", "PAWN", "PAWN", "PAWN", "PAWN", "PAWN", "PAWN"},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"pawn", "pawn", "pawn", "pawn", "pawn", "pawn", "pawn", "pawn"},
		{"rook", "knight", "bishop", "queen", "king", "bishop", "knight", "rook"}}

	return json.Marshal(board)
}

func PrintBoard(state []byte) (err error) {
	tempState := make([][]string, 64)
	err = json.Unmarshal(state, &tempState)
	if err != nil {
		return err
	}
	fmt.Printf("+ \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t +\n")
	fmt.Printf("| \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t |\n", tempState[0][0], tempState[0][1], tempState[0][2], tempState[0][3], tempState[0][4], tempState[0][5], tempState[0][6], tempState[0][7])
	fmt.Printf("+ \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t +\n")
	fmt.Printf("| \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t |\n", tempState[1][0], tempState[1][1], tempState[1][2], tempState[1][3], tempState[1][4], tempState[1][5], tempState[1][6], tempState[1][7])
	fmt.Printf("+ \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t +\n")
	fmt.Printf("| \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t |\n", tempState[2][0], tempState[2][1], tempState[2][2], tempState[2][3], tempState[2][4], tempState[2][5], tempState[2][6], tempState[2][7])
	fmt.Printf("+ \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t +\n")
	fmt.Printf("| \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t |\n", tempState[3][0], tempState[3][1], tempState[3][2], tempState[3][3], tempState[3][4], tempState[3][5], tempState[3][6], tempState[3][7])
	fmt.Printf("+ \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t +\n")
	fmt.Printf("| \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t |\n", tempState[4][0], tempState[4][1], tempState[4][2], tempState[4][3], tempState[4][4], tempState[4][5], tempState[4][6], tempState[4][7])
	fmt.Printf("+ \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t +\n")
	fmt.Printf("| \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t |\n", tempState[5][0], tempState[5][1], tempState[5][2], tempState[5][3], tempState[5][4], tempState[5][5], tempState[5][6], tempState[5][7])
	fmt.Printf("+ \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t +\n")
	fmt.Printf("| \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t |\n", tempState[6][0], tempState[6][1], tempState[6][2], tempState[6][3], tempState[6][4], tempState[6][5], tempState[6][6], tempState[6][7])
	fmt.Printf("+ \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t +\n")
	fmt.Printf("| \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t | \t%v\t |\n", tempState[7][0], tempState[7][1], tempState[7][2], tempState[7][3], tempState[7][4], tempState[7][5], tempState[7][6], tempState[7][7])
	fmt.Printf("+ \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t + \t------\t +\n")

	return err
}

// 用于判断移动是否合法
func IsValidMove(move model.Move, board model.Board) (isValid bool, err error) {
	tempState := make([][]string, 64)
	err = json.Unmarshal(board.State, &tempState)
	if err != nil {
		return false, err
	}

	// 解析 From 和 To 字段
	fromX := int(move.From[0] - 'a')
	fromY := int(move.From[1] - '0')
	toX := int(move.To[0] - 'a')
	toY := int(move.To[1] - '0')

	log.Printf("moveID: %v\nwhiteID: %v\nblackID: %v\n", move.UserID, board.BlackID)

	log.Printf("from (%v, %v) to (%v, %v)\n", fromX, fromY, toX, toY)

	// 起止位置不能相同
	if fromX == toX && fromY == toY {
		return false, nil
	}

	// 检查坐标是否在棋盘范围内
	if fromX < 0 || fromX > 7 || fromY < 0 || fromY > 7 || toX < 0 || toX > 7 || toY < 0 || toY > 7 {
		return false, nil
	}

	// 获取棋盘上的棋子
	fromPiece := tempState[fromY][fromX]
	toPiece := tempState[toY][toX]

	log.Printf("piece: %v\n", fromPiece)

	// 本轮只能移动轮次对应的白棋或黑棋
	if board.Turn == true { // 轮到白棋
		if move.UserID == board.BlackID {
			return false, nil
		}
	} else if board.Turn == false { // 轮到黑棋
		if move.UserID == board.WhiteID {
			return false, nil
		}
	}

	// 白方只能移动白棋，黑方只能移动黑棋
	if move.UserID == board.WhiteID {
		if InSlice(BlackPieces, fromPiece) {
			return false, nil
		}
	} else if move.UserID == board.BlackID {
		if InSlice(WhitePieses, fromPiece) {
			return false, nil
		}
	} else if (move.UserID != board.BlackID) && (move.UserID != board.WhiteID) {
		return false, nil
	}

	// 不能走到相同颜色的棋子上
	if InSlice(WhitePieses, fromPiece) {
		if InSlice(WhitePieses, toPiece) {
			return false, nil
		}
	} else if InSlice(BlackPieces, fromPiece) {
		if InSlice(BlackPieces, toPiece) {
			return false, nil
		}
	}

	// 根据棋子类型判断移动是否合法
	switch fromPiece {
	case "PAWN": // 白色兵
		if fromX == toX && fromY-toY == -1 { //向前走一格
			return true, nil
		} else if fromX == toX && abs(fromY-toY) == 2 && fromY == 1 { //向前走两格且在初始位置
			for fromX != toX || fromY != toY { //遍历路径
				if fromX < toX {
					fromX++
				} else if fromX > toX {
					fromX--
				}
				if fromY < toY {
					fromY++
				} else if fromY > toY {
					fromY--
				}
				if InSlice(WhitePieses, tempState[toY][toX]) || InSlice(BlackPieces, tempState[toY][toX]) {
					return false, nil //沿途有棋子
				}
			}
			return true, nil
		} else if abs(fromX-toX) == 1 && fromY-toY == -1 && InSlice(BlackPieces, toPiece) { //向斜前方走一格并且吃掉对方棋子
			return true, nil
		} else {
			return false, nil
		}
	case "pawn": // 黑色兵
		if fromX == toX && fromY-toY == 1 { //向前走一格
			return true, nil
		} else if fromX == toX && abs(fromY-toY) == 2 && fromY == 6 { //向前走两格且在初始位置
			for fromX != toX || fromY != toY { //遍历路径
				if fromX < toX {
					fromX++
				} else if fromX > toX {
					fromX--
				}
				if fromY < toY {
					fromY++
				} else if fromY > toY {
					fromY--
				}
				if InSlice(WhitePieses, tempState[toY][toX]) || InSlice(BlackPieces, tempState[toY][toX]) {
					return false, nil //沿途有棋子
				}
			}
			return true, nil
		} else if abs(fromX-toX) == 1 && fromY-toY == 1 && InSlice(BlackPieces, toPiece) { //向斜前方走一格并且吃掉对方棋子
			return true, nil
		} else {
			return false, nil
		}
	case "ROOK": // 白色车
		if fromX == toX || fromY == toY { //保证垂直移动
			for fromX != toX || fromY != toY { //遍历路径
				if fromX < toX {
					fromX++
				} else if fromX > toX {
					fromX--
				}
				if fromY < toY {
					fromY++
				} else if fromY > toY {
					fromY--
				}
				if InSlice(WhitePieses, tempState[toY][toX]) || InSlice(BlackPieces, tempState[toY][toX]) {
					return false, nil //沿途有棋子
				}
			}
			return true, nil
		} else {
			return false, nil
		}
	case "rook": // 黑色车
		if fromX == toX || fromY == toY { //保证垂直移动
			for fromX != toX || fromY != toY { //遍历路径
				if fromX < toX {
					fromX++
				} else if fromX > toX {
					fromX--
				}
				if fromY < toY {
					fromY++
				} else if fromY > toY {
					fromY--
				}
				if InSlice(WhitePieses, tempState[toY][toX]) || InSlice(BlackPieces, tempState[toY][toX]) {
					return false, nil //沿途有棋子
				}
			}
			return true, nil
		} else {
			return false, nil
		}
	case "QUEEN": // 白色后
		if fromX == toY || fromY == toY || abs(fromX-toX) == abs(fromY-toY) { //保证垂直或斜向移动
			for fromX != toX || fromY != toY { //遍历路径
				if fromX < toX {
					fromX++
				} else if fromX > toX {
					fromX--
				}
				if fromY < toY {
					fromY++
				} else if fromY > toY {
					fromY--
				}
				if InSlice(WhitePieses, tempState[toY][toX]) || InSlice(BlackPieces, tempState[toY][toX]) {
					return false, nil //沿途有棋子
				}
			}
			return true, nil
		} else {
			return false, nil
		}
	case "queen": // 黑色后
		if fromX == toY || fromY == toY || abs(fromX-toX) == abs(fromY-toY) { //保证垂直或斜向移动
			for fromX != toX || fromY != toY { //遍历路径
				if fromX < toX {
					fromX++
				} else if fromX > toX {
					fromX--
				}
				if fromY < toY {
					fromY++
				} else if fromY > toY {
					fromY--
				}
				if InSlice(WhitePieses, tempState[toY][toX]) || InSlice(BlackPieces, tempState[toY][toX]) {
					return false, nil //沿途有棋子
				}
			}
			return true, nil
		} else {
			return false, nil
		}
	case "KNIGHT": // 白色马
		if (abs(fromX-toX) == 2 && abs(fromY-toY) == 1) || (abs(fromY-toY) == 2 && abs(fromX-toX) == 1) { //保证走“日”字
			return true, nil
		} else {
			return false, nil
		}
	case "knight": // 黑色马
		if (abs(fromX-toX) == 2 && abs(fromY-toY) == 1) || (abs(fromY-toY) == 2 && abs(fromX-toX) == 1) { //保证走“日”字
			return true, nil
		} else {
			return false, nil
		}
	case "BISHOP": // 白色象
		if abs(fromX-toX) == abs(fromY-toY) { //保证斜向移动
			for fromX != toX || fromY != toY { //遍历路径
				if fromX < toX {
					fromX++
				} else if fromX > toX {
					fromX--
				}
				if fromY < toY {
					fromY++
				} else if fromY > toY {
					fromY--
				}
				if InSlice(WhitePieses, tempState[toY][toX]) || InSlice(BlackPieces, tempState[toY][toX]) {
					return false, nil //沿途有棋子
				}
			}
			return true, nil
		} else {
			return false, nil
		}
	case "bishop": // 黑色象
		if abs(fromX-toX) == abs(fromY-toY) { //保证斜向移动
			for fromX != toX || fromY != toY { //遍历路径
				if fromX < toX {
					fromX++
				} else if fromX > toX {
					fromX--
				}
				if fromY < toY {
					fromY++
				} else if fromY > toY {
					fromY--
				}
				if InSlice(WhitePieses, tempState[toY][toX]) || InSlice(BlackPieces, tempState[toY][toX]) {
					return false, nil //沿途有棋子
				}
			}
			return true, nil
		} else {
			return false, nil
		}
	case "KING": // 白色王
		if abs(fromX-toX) <= 1 && abs(fromY-toY) <= 1 {
			return true, nil
		} else {
			return false, nil
		}
	case "king": // 黑色王
		if abs(fromX-toX) <= 1 && abs(fromY-toY) <= 1 {
			return true, nil
		} else {
			return false, nil
		}
	default:
		return false, nil
	}
}
