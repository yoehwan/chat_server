package api

import (
	controller "chat/controller"

	"github.com/gin-gonic/gin"
)

var userController *controller.UserController

func HandleUserApi(r *gin.Engine, c *controller.UserController) {
	userController = c
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
