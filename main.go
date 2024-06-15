package main

import (
	"log"
	"net/http"

	"github.com/devder/go_event_booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	err := server.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, gin.H{"events": events})
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

	event.Save()
	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
