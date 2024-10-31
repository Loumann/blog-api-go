package middleware

import (
	"blog-api-go/controller"
	"blog-api-go/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (c controller.Controller) GenerateJWT(userId int) (string, error) {
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

func (c controller.Controller) ParserJWT(context *gin.Context, claims *models.Claims) {
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
	fmt.Printf(strconv.Itoa(claims.UserId))
	if claims.UserId == 0 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
}
