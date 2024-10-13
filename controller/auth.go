package controller

import (
	"blog-api-go/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var jwtSecret = []byte("my_secret_key")

func (c Controller) GetUsers(context *gin.Context) {

	users, err := c.Services.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal("error getting users", err.Error())
	}
	context.JSON(200, gin.H{"profile": users})
	context.AbortWithStatus(http.StatusOK)
}
func (c Controller) SignUp(context *gin.Context) {

	var input models.User
	var hashPass models.Credentials
	pass := []byte(input.Password)

	if err := context.ShouldBindJSON(&input); err != nil {
		log.Println("Error parsing body:", err)
	}

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
		log.Println("Password mismatch:", errs)
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

func (c Controller) GetProfile(context *gin.Context) {
	tokenString, err := context.Cookie("token")
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Парсинг токена и валидация
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil // Ваш секретный ключ
	})

	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	fmt.Printf(strconv.Itoa(claims.UserId))
	// Проверка наличия userId в токене
	if claims.UserId == 0 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	// Получение профиля пользователя из базы данных
	user, _, err := c.Services.GetProfileUser(claims.UserId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Возвращаем профиль пользователя
	context.JSON(http.StatusOK, gin.H{"profile": user})
}
