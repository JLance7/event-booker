package routes

import (
	"api/models"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	fmt.Println("Getting events")
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}
	// fmt.Println(id)
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success", "event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data. Try again later."})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save() // also sets event id
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could retrieve event."})
		return
	}
	if event.UserID != userId { // authorization, does event's userID == the userId of user sending request (from their JWT)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent) // place user's submitted values in updatedEvent var
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data. Try again later."})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error updating event. Try again later."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}
	deleteEvent, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	userId := context.GetInt64("userId")
	if deleteEvent.UserID != userId { // authorization, does event's userID == the userId of user sending request (from their JWT)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	deleteEvent.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error deleting event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}


