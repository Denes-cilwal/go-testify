package controllers

import "go-testify/services"

type UserController struct {
	Userservice services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		Userservice: userService,
	}
}
