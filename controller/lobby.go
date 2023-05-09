package controller

import (
	model "chat/model"
	"fmt"
)

type LobbyController struct {
	roomList []*model.Room
	userList []*model.User
}

func NewLobbyController() *LobbyController {
	return &LobbyController{
		roomList: []*model.Room{},
	}
}

func (c *LobbyController) UpdateRoom(room *model.Room) {
	for index, item := range c.roomList {
		if room.Id == item.Id {
			c.roomList[index] = item
			return
		}
	}
	c.roomList = append(c.roomList, room)
}

func (c *LobbyController) RemoveRoom(id string) {
	roomList := c.roomList
	for index, item := range roomList {
		if item.Id == id {
			c.roomList = append(roomList[:index], roomList[index+1:]...)
		}
	}
}

func (c *LobbyController) LoadRoomList() []*model.Room {
	return c.roomList
}

func (c *LobbyController) LoadRoom(id string) (*model.Room, error) {
	for _, item := range c.roomList {
		if item.Id == id {
			return item, nil
		}
	}
	return nil, fmt.Errorf("no Such Room ID %s ", id)
}

func (c *LobbyController) EnterRoom(roomID, userID string) error {
	room, err := c.LoadRoom(roomID)
	if err != nil {
		return err
	}
	user, err := c.LoadUser(userID)
	if err != nil {
		return err
	}
	room.UserList = append(room.UserList, user)
	return nil
}

func (c *LobbyController) LoadUser(userID string) (*model.User, error) {
	for _, item := range NewLobbyController().userList {
		if userID == item.Id {
			return item, nil
		}
	}
	return nil, fmt.Errorf("no such User ID %s", userID)
}

func (c *LobbyController) UpdateUser(user *model.User) {
	for index, item := range c.userList {
		if user.Id == item.Id {
			c.userList[index] = user
			return
		}
	}
	c.userList = append(c.userList, user)
}
func (c *LobbyController) RemoveUser(id string) {
	userList := c.userList
	for index, item := range userList {
		if item.Id == id {
			c.userList = append(userList[:index], userList[index+1:]...)
		}
	}
}

func (c *LobbyController) LoadUserList() []*model.User {
	return c.userList
}
