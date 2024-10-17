package controller

import (
	"blog-api-go/service"
)

type Controller struct {
	Services service.Service
}

func NewController(userService *service.Services) *Controller {
	return &Controller{Services: userService}
}
