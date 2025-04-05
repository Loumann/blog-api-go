package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func (c *Controller) InitRouters() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Разрешенные домены
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешенные HTTP-методы
		AllowHeaders:     []string{"Content-Type", "Authorization"},           // Разрешенные заголовки
		ExposeHeaders:    []string{"Content-Length"},                          // Заголовки, доступные клиенту
		AllowCredentials: true,                                                // Разрешить отправку cookies
		MaxAge:           12 * time.Hour,                                      // Кэширование CORS настроек
	}))

	router.LoadHTMLGlob("static/template/*")
	router.Static("/static", "./static")

	router.GET("/", c.LoginPage)
	router.GET("/feed", c.OnPointWindowLocation)
	router.GET("/registration", c.SignUpPage)
	router.GET("/users/:login", c.GetProfileFromLogin)
	router.GET("/page", c.MyPage)
	router.GET("/search", c.SearchPage)

	router.GET("/users", c.GetProfile)

	post := router.Group("/post")
	{
		post.GET("/", c.GetPosts)
		post.POST("/", c.CreatePost)
		post.DELETE("/:postId", c.DeletePost)
		post.PUT("/", c.UpdatePost)
	}

	subscribe := router.Group("/subscribe")
	{
		subscribe.POST("/:userId", c.Subscribe)
		//subscribe.DELETE("/:userId", c.Unsubscribe)
	}

	comment := router.Group("/comment")
	{
		comment.POST("/", c.CreateComment)
		comment.GET("/", c.GetComments)
		comment.DELETE("/:commentId", c.DeleteComment)
		comment.PUT("/", c.ChangeComment)
	}

	router.POST("/sign-in", c.SignIn)
	router.POST("/sign-up", c.SignUp)

	return router
}
