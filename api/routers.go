package api

import (
	"Gone/dao"
	"Gone/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	chess := r.Group("/chess")
	{
		chess.POST("/create", AuthMiddleware(), CreateBoard(service.H, dao.DB))
		chess.POST("/join", AuthMiddleware(), JoinBoard(service.H, dao.DB))
		chess.GET("/board/:boardid", AuthMiddleware(), BoardWebSocket(service.H, dao.DB))
	}

	user := r.Group("/user")
	{
		user.POST("/register", Register(service.H, dao.DB))
		user.POST("/login", Login(service.H, dao.DB))
	}

	r.Run(":9099")
}
