package main

import (
	"log"

	"github.com/devder/go_event_booking/db"
	"github.com/devder/go_event_booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	err := server.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
