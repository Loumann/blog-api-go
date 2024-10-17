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
	router.GET("/feed", c.OnPointWindowLocation)
	router.GET("/registration", c.SignUpPage)

	router.GET("/users", c.GetProfile)

	router.GET("/posts", c.GetPosts)

	post := router.Group("/post")
	{
		post.POST("/create", c.CreatePost)
		post.GET("/")
	}

	router.POST("/create-com", c.CreateComment)
	router.GET("/comments", c.GetComments)

	router.POST("/sig-in", c.SignIn)
	router.POST("/sig-up", c.SignUp)

	return router
}
