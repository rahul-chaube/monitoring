package router

import (
	"github.com/gin-gonic/gin"

	"github.com/rahul-chaube/monitoring/common"
	"github.com/rahul-chaube/monitoring/eventService/event"
	"github.com/rahul-chaube/monitoring/eventService/handler"
	"github.com/rahul-chaube/monitoring/eventService/repository"
	"github.com/rahul-chaube/monitoring/notificationService"
	"github.com/rahul-chaube/monitoring/uploader"
)

func EventRoute(r *gin.Engine) {
	s3Uploader := uploader.NewS3Uploader("mytestingbucket0406")
	eventMongoClient := common.MongoConnect("EventRepository")
	eventRepo := repository.NewEventRepository(eventMongoClient)
	eventService := event.NewEventService(eventRepo)
	notification := notificationService.NewNotificationService()
	eventHandler := handler.NewEventHandler(eventService, s3Uploader, notification)

	eventGroup := r.Group("/event")
	{
		eventGroup.POST("", eventHandler.AddEvent)
		eventGroup.GET("/list", eventHandler.ListEvent)
		eventGroup.GET("/notification/:token", eventHandler.SendNotication)
		eventGroup.GET("/:id", eventHandler.GetEvent)

	}
}
