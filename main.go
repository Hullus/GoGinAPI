package main

import (
	"MaxRestAPI/db"
	"MaxRestAPI/routes"
	"errors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		errors.New("")
	}
}
