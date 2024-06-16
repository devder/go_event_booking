package routes

import (
	"fmt"
	"net/http"

	"github.com/devder/go_event_booking/models"
	"github.com/devder/go_event_booking/utils"
	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "successfully signed up"})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token for user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logged in", "token": token})

}
