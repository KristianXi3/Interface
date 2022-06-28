package main

import (
	"INTERFACE/entity1"
	"INTERFACE/service"
	"time"
)

func main() {

	userSvc := service.NewUserService()

	userSvc.Register(&entity1.User{
		Id:        1,
		Username:  "Kristian",
		Email:     "email@email.com",
		Password:  "Password",
		Age:       17,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

}
