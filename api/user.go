package api

import (
	controller "chat/controller"

	"github.com/gin-gonic/gin"
)

func HandleUserApi(r *gin.Engine, c *controller.LobbyController) {
	updateUser(r)
	loadUserData(r)
	removeUser(r)
	loadUserList(r)
}

func updateUser(r *gin.Engine) {

}

func loadUserData(r *gin.Engine) {

}

func removeUser(r *gin.Engine) {

}

func loadUserList(r *gin.Engine) {
}
