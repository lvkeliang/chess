package model

import (
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"log"
)

// Hub 是一个 websocket 中心，它维护了活跃的连接和广播消息
type Hub struct {
	Clients     map[uint]map[*websocket.Conn]bool // 一个存储活跃 websocket 连接的集合
	Broadcast   chan []byte                       // 一个用于向所有 Clients 广播消息的通道
	UpdateBoard chan Board
	Register    chan map[uint]*websocket.Conn // 一个用于注册新 Clients 的通道
	Unregister  chan map[uint]*websocket.Conn // 一个用于注销已有 Clients 的通道
}

// Run 在一个无限循环中运行 hub，处理不同的通道
func (h *Hub) Run() {
	for {
		select {
		case register := <-h.Register: // 当有新的连接被注册时
			for boardID, conn := range register { // 得到房间ID和连接
				if _, ok := h.Clients[boardID]; !ok { // 如果房间不存在
					h.Clients[boardID] = make(map[*websocket.Conn]bool) // 创建一个新的房间
				}
				h.Clients[boardID][conn] = true // 将连接加入到房间中
			}

		case unregister := <-h.Unregister: // 当有已有的连接被注销时
			for boardID, conn := range unregister { // 得到房间ID和连接
				if _, ok := h.Clients[boardID][conn]; ok { // 如果连接在房间中
					delete(h.Clients[boardID], conn) // 将连接从房间中删除
					conn.Close()                     // 关闭连接
				}
			}

		case message := <-h.Broadcast: // 当有新的消息被广播时
			for boardID := range h.Clients { // 对于每个房间
				for conn := range h.Clients[boardID] { // 对于每个连接
					conn.WriteMessage(websocket.TextMessage, message) // 将消息写入到连接中
				}
			}
		case board := <-h.UpdateBoard: // 当有新的消息被广播时
			for conn := range h.Clients[board.ID] { // 对于每个连接
				message, err := json.Marshal(board) // 将消息编码为 JSON 格式
				if err != nil {
					log.Println(err)
					continue
				}
				conn.WriteMessage(websocket.TextMessage, message) // 将消息写入到连接中
			}
		}
	}
}
