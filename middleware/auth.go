package middleware

import (
	"gin-api/internal/authentication"
	"gin-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing auth header"})
	}

	userId, err := authentication.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate the token", "error": err.Error()})
	}

	if userId == 0 {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Token has no user ID"})
	}

	_, err = models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "Auth user does not exist"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
