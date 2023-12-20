package routes

import (
	"MaxRestAPI/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authanticate)

	//Event routes
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)

	//Authenticated
	authenticated.POST("/events", CreateEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

	//User routes
	server.POST("/signup", SignUp)
	server.POST("/login", Login)
}
