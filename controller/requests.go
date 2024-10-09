package controller

import (
	"github.com/gin-gonic/gin"
)

func (c *Controller) InitRouters() *gin.Engine {
	router := gin.Default()

	authorization := router.Group("/sig-in")
	{
		authorization.GET("/sig-in")

	}

	router.LoadHTMLGlob("template/*")
	router.Static("/static", "./static")

	router.GET("/", c.LoginPage)
	router.GET("/users", c.GetUsers)
	router.GET("feed", c.OnPointWindowLocation)

	router.POST("/sig-in", c.SignIn)

	return router
}
