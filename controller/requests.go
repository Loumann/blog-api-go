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

	post := router.Group("/post")
	{
		post.GET("/", c.GetPosts)
		post.POST("/create", c.CreatePost)
		post.DELETE("/delete/:postId", c.DeletePost)
		post.PUT("/change", c.ChangePost)
	}

	comment := router.Group("/comment")
	{
		comment.POST("/create-com", c.CreateComment)
		comment.GET("/", c.GetComments)
		comment.DELETE("/delete/:commentId", c.DeleteComment)
		comment.PUT("/change", c.ChangeComment)
	}

	router.POST("/sig-in", c.SignIn)
	router.POST("/sig-up", c.SignUp)

	return router
}
