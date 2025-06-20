package controller

import (
	"github.com/gin-gonic/gin"
)

func (c *Controller) InitRouters() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("static/template/*")
	router.Static("/static", "./static")

	router.GET("/", c.LoginPage)
	router.GET("/feed", c.OnPointWindowLocation)
	router.GET("/registration", c.SignUpPage)
	router.GET("/users/:login", c.GetProfileFromLogin)
	router.GET("/page", c.MyPage)
	router.GET("/search", c.SearchPage)
	router.GET("/subs", c.SubsPages)

	router.GET("/users", c.GetProfile)
	router.GET("/check-sub/:userId", c.CheckSubscribe)

	post := router.Group("/post")
	{
		post.POST("/", c.CreatePost)
		post.GET("/", c.GetPosts)
		post.DELETE("/:postId", c.DeletePost)
		post.PUT("/", c.UpdatePost)
	}

	subscribe := router.Group("/subscribe")
	{
		subscribe.GET("/posts", c.GetPostsSubcribe)
		subscribe.POST("/:userId", c.Subscribe)
	}

	comment := router.Group("/comment")
	{
		comment.POST("/:postId", c.CreateComment)
		comment.GET("/", c.GetComments)
		comment.DELETE("/:commentId", c.DeleteComment)
		comment.PUT("/", c.ChangeComment)
	}

	router.POST("/sign-in", c.SignIn)
	router.POST("/sign-up", c.SignUp)

	return router
}
