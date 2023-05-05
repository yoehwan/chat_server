package controller

import model "chat/model"

type UserController struct {
	userList []*model.User
}

func (c *UserController) UpdateUser(user *model.User) {
	c.userList = append(c.userList, user)
}
func (c *UserController) RemoveUser(id int) {
	userList := c.userList
	for index, item := range userList {
		if item.Id == id {
			c.userList = append(userList[:index], userList[index+1:]...)
		}
	}
}

func (c *UserController) LoadUserList() []*model.User {
	return c.userList
}
