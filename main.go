package main

import (
	"log"

	"github.com/devder/go_event_booking/db"
	"github.com/devder/go_event_booking/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDB()
	server := gin.Default()

	err = server.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
