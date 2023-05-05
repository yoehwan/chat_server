package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoomApi(r *gin.Engine) {
	updateRoom(r)
	loadRoomData(r)
	removeRoom(r)
	loadRoomList(r)
	enterRoom(r)
	leaveRoom(r)
}

func loadRoomData(r *gin.Engine) {
	r.GET("/room", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})
}

func removeRoom(r *gin.Engine) {
	r.DELETE("/room", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})
}

func updateRoom(r *gin.Engine) {
	r.POST("/room", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
}

func enterRoom(r *gin.Engine) {
	r.GET("/enterRoom", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
}
func leaveRoom(r *gin.Engine) {
	r.GET("/leaveRoom", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
}

func loadRoomList(r *gin.Engine) {
	r.GET("/rooms", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})
}
