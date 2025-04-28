package controller

import (
	"blog-api-go/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (c Controller) GenerateJWT(userId int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (c Controller) ParserJWT(context *gin.Context, claims *models.Claims) error {
	tokenString, err := context.Cookie("token")
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return err
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return err
	}

	if claims.UserId == 0 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return err
	}
	return err
}
