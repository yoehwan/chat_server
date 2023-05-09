package model

import (
	"fmt"
	"time"
)

type User struct {
	Id   string
	Name string
}

func NewUser(name string) *User {
	return &User{
		Id:   fmt.Sprintf("%d", time.Now().UnixMilli()),
		Name: name,
	}
}
