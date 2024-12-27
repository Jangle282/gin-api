package routes

import (
	"gin-api/handlers"
	"gin-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// getEvents createEvents are handlers

	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:event_id", handlers.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	authenticated.POST("/events", handlers.CreateEvent)
	authenticated.PUT("/events/:event_id", handlers.UpdateEvent)
	authenticated.DELETE("/events/:event_id", handlers.DeleteEvent)
	authenticated.POST("/events/:event_id/users/:user_id", handlers.ConfirmEvent)
	authenticated.DELETE("/events/:event_id/users/:user_id", handlers.DeclineEvent)

	server.POST("/users", handlers.CreateUser)
	server.POST("/login", handlers.Login)
}
