package service

import (
	"Gone/model"
	"github.com/gorilla/websocket"
)

var H *model.Hub

// newHub creates a new hub instance
func NewHub() *model.Hub {
	h := &model.Hub{
		Clients:     make(map[uint]map[*websocket.Conn]bool),
		Broadcast:   make(chan []byte),
		UpdateBoard: make(chan model.Board),
		Register:    make(chan map[uint]*websocket.Conn),
		Unregister:  make(chan map[uint]*websocket.Conn),
	}

	H = h

	return h
}
