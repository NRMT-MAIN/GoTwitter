package controllers

import "GoTwitter/m/service"

type UserController struct {
	userService *service.UserService
}

func NewUserController(_userService *service.UserService) {
	return &UserController{
		userService : _userService
	}
}


