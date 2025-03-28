package controller

import (
	"blog-api-go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (c Controller) ChangeComment(context *gin.Context) {
	var comment models.Comments
	err := context.ShouldBindJSON(&comment)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	update, err := c.Services.ChangeComment(comment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	if !update {
		context.JSON(http.StatusBadRequest, gin.H{"error": "comment not exist"})
	}
	context.JSON(http.StatusOK, gin.H{"status": "comment updated"})

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
		context.AbortWithStatusJSON(500, gin.H{"error": "post not exist"})
		return
	}

	com, err := c.Services.CreateComment(claims.UserId, postId, input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"create new comment": com})

}

func (c *Controller) GetComments(context *gin.Context) {
	post, err := c.Services.GetComments()
	if err != nil {
		context.JSON(401, gin.H{"error": err})
		log.Fatal("error getting posts", err.Error())
	}
	context.AbortWithStatusJSON(http.StatusOK, post)
	return
}

func (c Controller) DeleteComment(context *gin.Context) {
	commentIdParam := context.Param("commentId")
	commentId, err := strconv.Atoi(commentIdParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
	}
	err = c.Services.DeleteComment(commentId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"status": "comment deleted"})
}
