package handler

import (
	"Monitoring/eventService/event"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	event event.EventService
}

func NewEventHandler(event event.EventService) *EventHandler {
	return &EventHandler{event: event}
}

func (h *EventHandler) AddEvent(c *gin.Context) {

	c.JSON(200, gin.H{
		"event": "thanks",
	})
}
