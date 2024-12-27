package main

import (
	"gin-api/database"
	"gin-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialises the database
	err := db.InitDB()
	if err != nil {
		return
	}

	// ses up a basic server
	server := gin.Default()

	routes.RegisterRoutes(server)

	// listen to incoming requests on a given domain and port
	server.Run(":8080")
}
