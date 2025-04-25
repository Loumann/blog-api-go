package controller

import (
	"blog-api-go/internal/service"
)

type Controller struct {
	Services *service.Services
}

func NewController(userService *service.Services) *Controller {
	return &Controller{Services: userService}
}
