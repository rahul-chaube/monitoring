package handler

import (
	"Monitoring/eventService/event"
	"Monitoring/eventService/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type EventHandler struct {
	event event.EventService
}

func NewEventHandler(event event.EventService) *EventHandler {
	return &EventHandler{event: event}
}

func (h *EventHandler) AddEvent(c *gin.Context) {
	log.Println("Handler Add Event called ")
	_, err := h.event.AddEvent(model.Event{
		EventName: "TestEvent",
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"event": "thanks",
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
