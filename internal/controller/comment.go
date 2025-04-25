package controller

import (
	models2 "blog-api-go/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c Controller) ChangeComment(context *gin.Context) {
	var comment models2.Comments
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

func (c Controller) CreateComment(context *gin.Context) {
	var input models2.Comments

	postIDKey := context.Param("postId")
	id, err := strconv.Atoi(postIDKey)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	claims := &models2.Claims{}
	c.ParserJWT(context, claims)

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	com, err := c.Services.CreateComment(claims.UserId, id, input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"create new comment": com})

}

func (c Controller) GetComments(context *gin.Context) {
	postIdStr := context.Query("post_id")
	idPost, err := strconv.Atoi(postIdStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный post_id"})
		return
	}

	comments, err := c.Services.GetComments(idPost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if comments == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "comments not exist"})
		return
	}

	context.JSON(http.StatusOK, comments)
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
