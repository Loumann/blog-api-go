package controller

import (
	"blog-api-go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	has, err := c.Services.SignIn(input.Login)
	if err != nil {
		log.Println("Error signing in:", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid login or password"})
		return
	}

	inputPassword := input.Password
	errs := bcrypt.CompareHashAndPassword(has, []byte(inputPassword))
	if errs != nil {
		log.Println("Password mismatch:", errs)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid login or password"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})

}

func (c *Controller) SignUp(context *gin.Context) {
	var input models.User
	var hashPass models.Credentials

	if err := context.ShouldBindJSON(&input); err != nil {
		log.Println("Error parsing body:", err)
	}

	pass := []byte(input.Password)

	pas, _ := bcrypt.GenerateFromPassword(pass, 10)
	hashPass.Password = string(pas)

	if err := context.ShouldBindJSON(&hashPass); err != nil {
		log.Println("Error parsing body:", err)
	}

	err := c.Services.SignUp(input, hashPass)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal("error signing up:", err)

	} else {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "user added successfully"})
	}

}
