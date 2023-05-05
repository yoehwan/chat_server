package main

import (
	api "chat/api"
	controller "chat/controller"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

var userController *controller.UserController

func main() {
	userController = &controller.UserController{}
	r := gin.Default()
	r.Use(logger.SetLogger())
	api.HandleRoomApi(r)
	api.HandleUserApi(r, userController)

	r.Run("localhost:4330") // li
}
