package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahul-chaube/monitoring/eventService/event"
	"github.com/rahul-chaube/monitoring/eventService/model"
)

type EventHandler struct {
	event        event.EventService
	s3Upload     *uploader.S3Uploader
	notification *notificationService.NotificationService
}

func NewEventHandler(event event.EventService, s3 *uploader.S3Uploader, service *notificationService.NotificationService) *EventHandler {
	return &EventHandler{event: event, s3Upload: s3, notification: service}
}

func (h *EventHandler) AddEvent(c *gin.Context) {
	log.Println("Handler Add Event called ")
	eventName := c.PostForm("eventName")
	eventType := c.PostForm("eventType")
	var event model.Event
	event.EventName = eventName
	event.EventType = eventType

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}
	uploadedFiles := form.File["files"]
	var savedFiles []string

	for _, file := range uploadedFiles {
		//path := "uploads/" + file.Filename
		//if err := c.SaveUploadedFile(file, path); err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: " + file.Filename})
		//	return
		//}
		key, path, err := h.s3Upload.UploadFile(file)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		}
		fmt.Println(key, path)
		savedFiles = append(savedFiles, key)
	}

	signedUrl := h.s3Upload.Presigned(savedFiles)
	fmt.Println("All signed Url ", signedUrl)

	err = h.notification.SendMessage("cUIzEXTZXWN5p6P3-44Ywn:APA91bHsnCO-UzKygk0BI9EO9ngUSkQBFIAI8jJZY2Ydl0mQkf6g3YpfEkJmpho3KTMAfwrMcI8CnOd8a3zZLEGczG6iDzY3t-cVUDDvfhNuej3mo0Mqmgk", "Hello", "world")
	if err != nil {
		log.Println(err)
	}
	event.CreatedAt = time.Now().UTC()
	event.Files = savedFiles

	ev, err := h.event.AddEvent(event)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"message": "Event added successfully",
		"files":   savedFiles,
		"data":    ev,
	})
}

func (h *EventHandler) ListEvent(c *gin.Context) {
	log.Println("Handler All Event called ")
	events, err := h.event.GetAllEvents()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, events)
}
