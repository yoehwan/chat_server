package controller

import model "chat/model"

type RoomController struct {
	roomList []*model.Room
}

func (c *RoomController) UpdateRoom(room *model.Room) {
	c.roomList = append(c.roomList, room)
}

func (c *RoomController) RemoveRoom(id int) {
	roomList := c.roomList
	for index, item := range roomList {
		if item.Id == id {
			c.roomList = append(roomList[:index], roomList[index+1:]...)
		}
	}
}

func (c *RoomController) LoadRoomList() []*model.Room {
	return c.roomList
}
