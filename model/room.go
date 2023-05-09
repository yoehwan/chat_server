package model

import (
	"fmt"
	"time"
)

type Room struct {
	Id       string
	Name     string
	UserList []*User
}

func NewRoom(name string) *Room {
	return &Room{
		Id:       fmt.Sprintf("%d", time.Now().UnixMilli()),
		Name:     name,
		UserList: []*User{},
	}
}
