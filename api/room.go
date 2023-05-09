package api

import (
	controller "chat/controller"
	model "chat/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var lobbyController *controller.LobbyController

func HandleRoomApi(r *gin.Engine, c *controller.LobbyController) {
	lobbyController = c
	updateRoom(r)
	loadRoomData(r)
	removeRoom(r)
	loadRoomList(r)
	enterRoom(r)
	leaveRoom(r)
}

func loadRoomData(r *gin.Engine) {
	r.GET("/room/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var res *model.Response
		room, err := lobbyController.LoadRoom(id)
		if err != nil {
			res = model.FailResponse(
				err.Error(),
			)
		} else {
			res = model.SuccessResponse(room)
		}
		c.JSON(http.StatusOK, res)
	})
}

func removeRoom(r *gin.Engine) {
	r.DELETE("/room/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		lobbyController.RemoveRoom(id)
		c.JSON(http.StatusOK, model.SuccessResponse(nil))
	})
}

func enterRoom(r *gin.Engine) {
	r.GET("/enterRoom/:roomID/:userID", func(c *gin.Context) {
		roomID := c.Params.ByName("roomID")
		userID := c.Params.ByName("userID")
		err := lobbyController.EnterRoom(roomID, userID)
		var res *model.Response
		if err != nil {
			res = model.FailResponse(err.Error())
		} else {
			res = model.SuccessResponse(nil)
		}
		c.JSON(http.StatusOK, res)
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

func updateRoom(r *gin.Engine) {
	r.POST("/room", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
}
