package routes

import (
	"MaxRestAPI/models"
	"MaxRestAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"massage": "Could not parse request data.", "err": err})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massage:": "Could create user. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"massage": "User created"})
}

func Login(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"massage": "Could not parse request data.", "err": err})
		return
	}

	err = user.ValidateCredantials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"massage": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massage": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"massage": "Login success.", "token": token})
}
