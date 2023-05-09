package main

import (
	api "chat/api"
	controller "chat/controller"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

var lobbyController *controller.LobbyController

func main() {
	lobbyController = controller.NewLobbyController()
	r := gin.Default()
	r.Use(logger.SetLogger())
	api.HandleRoomApi(r, lobbyController)
	api.HandleUserApi(r, lobbyController)
	r.Run("localhost:4330")
}
