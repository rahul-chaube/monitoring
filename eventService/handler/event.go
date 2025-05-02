package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahul-chaube/monitoring/eventService/event"
	"github.com/rahul-chaube/monitoring/eventService/model"
	"github.com/rahul-chaube/monitoring/notificationService"
	"github.com/rahul-chaube/monitoring/uploader"
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
	var event model.Event
	if err := c.ShouldBind(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if event.Confidence < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "confidence is zero"})
		return
	}
	// Parse EventTime
	parsedEventTime, err := time.Parse(time.RFC3339, event.EventTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid eventTime format. Expected RFC3339 format (e.g., 2025-04-28T15:04:05Z)"})
		return
	}
	event.EventTime = parsedEventTime

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}
	uploadedFiles := form.File["files"]
	var savedFiles []string

	for _, file := range uploadedFiles {
		key, path, err := h.s3Upload.UploadFile(file)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
			return
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

	fmt.Printf("%+v", ev)
	ev.Files = signedUrl

	messageId := notificationService.SendSMS()
	fmt.Println("Message ID:", messageId)

	c.JSON(200, gin.H{
		"message": "Event added successfully",
		"data":    ev,
	})
	return
}

func (h *EventHandler) ListEvent(c *gin.Context) {
	log.Println("Handler All Event called ")
	events, err := h.event.GetAllEvents()

	for i := 0; i < len(events); i++ {
		events[i] = h.UpdatePresignedUrl(events[i])
	}
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, events)
}

func (h *EventHandler) GetEvent(c *gin.Context) {
	log.Println("Handler All Event called ")

	eventId := c.Param("id")

	if eventId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No event id"})
		return
	}
	event, err := h.event.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	event = h.UpdatePresignedUrl(event)
	c.JSON(200, event)

}

func (h *EventHandler) UpdatePresignedUrl(event model.Event) model.Event {
	event.Files = h.s3Upload.Presigned(event.Files)
	return event
}

func (h *EventHandler) SendNotication(c *gin.Context) {
	log.Println("Send Test Notification called ")

	token := c.Param("token")

	err := h.notification.SendMessage(token, "Test", "Test"+time.Now().String())
	if err != nil {

	}
	c.JSON(200, gin.H{
		"message": "Notification sent successfully",
	})

}
