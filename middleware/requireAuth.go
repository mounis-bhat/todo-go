package middleware

import (
	"example/todo-go/initializers"
	"example/todo-go/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var Id int64

func RequireAuth(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")

	fmt.Println(tokenString)

	if tokenString == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization header is missing",
		})
		return
	}

	accessToken := tokenString[len("Bearer "):]

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User

		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		Id = int64(user.ID)

		context.Next()
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}
