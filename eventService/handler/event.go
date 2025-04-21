package handler

import (
	"Monitoring/eventService/event"
	"Monitoring/eventService/model"
	"Monitoring/uploader"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type EventHandler struct {
	event    event.EventService
	s3Upload *uploader.S3Uploader
}

func NewEventHandler(event event.EventService, s3 *uploader.S3Uploader) *EventHandler {
	return &EventHandler{event: event, s3Upload: s3}
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
