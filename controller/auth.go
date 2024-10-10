package controller

import (
	"blog-api-go/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (c *Controller) GetUsers(context *gin.Context) {

	users, err := c.Services.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal("error getting users", err.Error())
	}
	context.JSON(200, users)
	context.AbortWithStatus(http.StatusOK)
}
func (c *Controller) SignIn(context *gin.Context) {

	var input models.User

	if err := context.ShouldBindJSON(&input); err != nil {
		log.Println("Error parsing body:", err)
		context.JSON(http.StatusBadRequest, gin.H{"Error parsing(maybe the body is null)": err.Error()})
		return
	}

	err := c.Services.SignIn(input.Login, input.Password)

	if err != nil {
		log.Println("Error signing in:", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid login or password"})
		return
	} else {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
	}

}
