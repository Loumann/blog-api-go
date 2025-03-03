package controller

import (
	"blog-api-go/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

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

func (c Controller) GetPosts(context *gin.Context) {
	page := context.DefaultQuery("page", "1")
	limit := context.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	posts, err := c.Services.GetPosts(pageInt, limitInt)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error getting posts", err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (c Controller) DeletePost(context *gin.Context) {
	postIdParam := context.Param("postId")
	postId, err := strconv.Atoi(postIdParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	err = c.Services.DeletePost(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "post deleted"})
}

func (c Controller) ChangePost(context *gin.Context) {
	var post models.Post

	err := context.ShouldBindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	update, err := c.Services.ChangePost(post)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	if !update {
		context.JSON(http.StatusBadRequest, gin.H{"error": "post not exist"})

	} else {
		context.JSON(http.StatusOK, gin.H{"status": "post updated"})
	}
}

func (c Controller) ChechRools(context *gin.Context, userId int) {
	claims := &models.Claims{}
	c.ParserJWT(context, claims)

}
