package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	err := server.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	server.GET("/events", getEvents)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello from the API"})
}
