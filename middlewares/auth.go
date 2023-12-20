package middlewares

import (
	"MaxRestAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authanticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"massage": "not authorized"})
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"massage": "not authorized"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
