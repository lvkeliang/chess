package main

import (
	"Gone/api"
	"Gone/dao"
	"Gone/service"
)

func main() {
	dao.InitDB()          // initialize the database connection and migration
	h := service.NewHub() // create a new hub instance for websocket connections and messages
	go h.Run()            // run
	api.InitRouter()

}
