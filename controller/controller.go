package controller

import (
	"blog-api-go/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct {
	Services service.Service
}

func NewController(userService service.Service) *Controller {
	return &Controller{Services: userService}
}

func (c *Controller) GetPost(context *gin.Context) {
	post, err := c.Services.GetPosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatal("error getting posts", err.Error())
	}

	context.AbortWithStatusJSON(http.StatusOK, post)
	return
}
