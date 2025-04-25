package controller

import (
	"blog-api-go/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c Controller) Subscribe(context *gin.Context) {
	userId := context.Param("userId")
	claims := &models.Claims{}
	c.ParserJWT(context, claims)

	err := c.Services.ToggleSub(claims.UserId, userId)
	if err {
		context.JSON(http.StatusInternalServerError, gin.H{"status": err})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": err})
	}

}

func (c Controller) CheckSubscribe(ctx *gin.Context) {
	userIdParam := ctx.Param("userId")

	claims := &models.Claims{}
	if err := c.ParserJWT(ctx, claims); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	subscribed, err := c.Services.CheckIfSubscribed(claims.UserId, userIdParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"subscribed": subscribed})
}
