package routes

import (
	"MaxRestAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"massage": "Could not get event. Try again later."})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"massage": "Could not find event. Try again later."})
	}

	context.JSON(http.StatusOK, event)
}

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massage:": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"massage": "Could not parse request data.", "err": err})
		return
	}
	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massage:": "Could create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"massage": "Event created", "event": event})

}

func DeleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"massage": "Event not found. Try again later."})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"massage": "Could not find event. Try again later."})
	}

	err = event.DeleteByID()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massage:": "Event not found. Try again later."})
	}

	context.JSON(http.StatusOK, gin.H{"massage": "Event Deleted."})
}

func UpdateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"massage": "Event not found. Try again later."})
		return
	}

	err = models.UpdateEventByID(event, eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"massage:": "Event could not be updated. Try again later."})
	}

	context.JSON(http.StatusOK, gin.H{"massage": "Event updated.", "Event": event})
}
