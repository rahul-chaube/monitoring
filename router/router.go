package router

import (
	"Monitoring/eventService/event"
	"Monitoring/eventService/handler"
	"Monitoring/eventService/repository"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	eventRepo := repository.NewEventRepository()
	eventService := event.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	eventGroup := r.Group("/event")
	{
		eventGroup.POST("/", eventHandler.AddEvent)
	}
	return r
}
