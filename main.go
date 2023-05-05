package main

import (
	api "chat/api"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(logger.SetLogger())
	api.HandleRoomApi(r)

	r.Run("localhost:4330") // li
}
