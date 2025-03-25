package controller

import (
	"blog-api-go/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (c Controller) CreatePost(context *gin.Context) {
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
	userId := context.DefaultQuery("own", "false")
	own, err := strconv.ParseBool(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	tokenString, err := context.Cookie("token")
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		return
	}

	var claims models.Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	posts, err := c.Services.GetPosts(claims.UserId, pageInt, limitInt, own)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Error getting posts:", err.Error())
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

func (c Controller) UpdatePost(context *gin.Context) {
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
