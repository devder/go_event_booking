package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/devder/go_event_booking/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events, please try again later"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"events": events}) // returns null for an empty array
}

func getEventById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id, please try again later"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("could not get event with id %v, please try again later", id)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"event": event})

}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event, please try again later"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
