package routes

import (
	"api/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


func registerForEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	userId := context.GetInt64("userId")
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering for event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event!"})
}

func cancelRegistration(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	userId := context.GetInt64("userId")
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering for event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Canceled registration for event!"})
}