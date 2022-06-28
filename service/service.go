package service

import (
	"INTERFACE/entity1"
	"fmt"
)

type UserServiceIface interface {
	Register(user *entity1.User)
}

type UserSvc struct{}

func NewUserService() UserServiceIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *entity1.User) {
	fmt.Printf("%s \n%s \n%s \n%d \n%v \n%v \n", user.Username, user.Email, user.Password, user.Age, user.CreatedAt, user.UpdatedAt)
}
