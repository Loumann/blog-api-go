package controller

import (
	"blog-api-go/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var jwtSecret = []byte("my_secret_key")

func (c Controller) SaveAvatar(context *gin.Context) {

}

func (c Controller) SignUp(context *gin.Context) {

	var input models.User
	var hashPass models.Credentials

	if err := context.ShouldBindJSON(&input); err != nil {
		log.Println("Error parsing body:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	pass := []byte(input.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	hashPass.Password = string(hashedPass)

	err = c.Services.SignUp(input, hashPass)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal("Error signing up:", err)
	} else {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "User added successfully"})
	}

}
func (c Controller) SignIn(context *gin.Context) {
	var input models.User

	if err := context.ShouldBindJSON(&input); err != nil {
		log.Println("Error parsing body:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing body"})
		return
	}

	hashedPassword, userid, err := c.Services.SignIn(input.Login)
	if err != nil {
		log.Println("Error signing in:", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid login or password"})
		return
	}

	errs := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password))
	if errs != nil {
		log.Println(input.Password)
		log.Println("Password mismatch:", errs)
		log.Printf(string(hashedPassword))

		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid login or password"})
		return
	}

	jwToken, err := c.GenerateJWT(userid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT"})
		return
	}

	context.SetCookie("token", jwToken, 3600*24, "/", "localhost", true, true)
	context.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"token":   jwToken,
	})
}

func (c Controller) GetUsers(context *gin.Context) {

	users, err := c.Services.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal("error getting users", err.Error())
	}
	context.JSON(200, gin.H{"profile": users})
	context.AbortWithStatus(http.StatusOK)
}

func (c Controller) GetProfile(context *gin.Context) {
	claims := &models.Claims{}

	c.ParserJWT(context, claims)

	user, _, err := c.Services.GetProfileUser(claims.UserId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (c Controller) GetProfileFromLogin(context *gin.Context) {
	login := context.Param("login")
	user, _, err := c.Services.GetProfileUserForLogin(login)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)

}

func (c Controller) GenerateJWT(userId int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("my_secret_key")) // Замените на ваш секретный ключ
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (c Controller) ParserJWT(context *gin.Context, claims *models.Claims) {
	tokenString, err := context.Cookie("token")
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil // Ваш секретный ключ
	})

	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	if claims.UserId == 0 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
}
