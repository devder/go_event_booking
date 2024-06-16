package routes

import (
	"fmt"
	"net/http"

	"github.com/devder/go_event_booking/models"
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

	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}
