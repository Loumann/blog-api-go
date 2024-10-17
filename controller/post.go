package controller

import (
	"blog-api-go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (c *Controller) GetComments(context *gin.Context) {
	post, err := c.Services.GetComments()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatal("error getting posts", err.Error())
	}
	context.AbortWithStatusJSON(http.StatusOK, post)
	return
}

func (c *Controller) CreatePost(context *gin.Context) {
	var Post models.Post
	claims := &models.Claims{}
	c.ParserJWT(context, claims)
	if err := context.ShouldBindJSON(&Post); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := c.Services.CreatePost(claims.UserId, Post)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatal("error getting posts", err.Error())
	}
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"create new post": Post})
}

func (c *Controller) CreateComment(context *gin.Context) {
	var input models.Comments
	claims := &models.Claims{}

	c.ParserJWT(context, claims)
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		fmt.Printf(err.Error())
	}
	postId, err := c.Services.GetIdPost(input.Id_post_on_comment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatal("error getting posts ", err.Error())
	}

	com, err := c.Services.CreateComment(claims.UserId, postId, input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatal("error getting comments", err.Error())
	}
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"create new comment": com})

}

func (c Controller) GetPosts(context *gin.Context) {
	posts, err := c.Services.GetPosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatal("error getting posts", err.Error())
	}
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"posts": posts})
	return
}
