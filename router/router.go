package router

import (
	"Monitoring/common"
	"Monitoring/eventService/event"
	"Monitoring/eventService/handler"
	"Monitoring/eventService/repository"
	"Monitoring/uploader"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	s3Uploader := uploader.NewS3Uploader("mytestingbucket0406")
	eventMongoClient := common.MongoConnect("EventRepository")
	eventRepo := repository.NewEventRepository(eventMongoClient)
	eventService := event.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService, s3Uploader)

	eventGroup := r.Group("/event")
	{
		eventGroup.POST("", eventHandler.AddEvent)
		eventGroup.GET("/list", eventHandler.ListEvent)
	}
	return r
}
