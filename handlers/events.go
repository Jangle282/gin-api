package handlers

import (
	"gin-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Couldn't fetch events": err.Error()})
		return
	}
	//can access the request / make a response with context
	//request := context.Request.Body
	//dont return.
	// object can be string, number, most often a struct or obkect
	// gin has a custom type H help to make a map to send back that can have any value.
	context.JSON(http.StatusOK, gin.H{
		"message": "here are the events",
		"data":    events,
	})
}

// context passed if createEVent is registered as a handler - e.g by passing it to a route above
func CreateEvent(context *gin.Context) {
	// validation don't with struct tags on the model struct
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"couldnt parse request data": err.Error()})
		// need to return otherwise rest of function would execute
		return
	}

	event.UserId = context.GetInt64("userId")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Couldn't create events": err.Error()})
		return
	}

	// don't return anything - use the context object
	context.JSON(http.StatusCreated, gin.H{
		"message": "event created",
		"data":    event,
	})
}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("event_id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"couldn't parse event ID": err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"couldn't get event by ID": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("event_id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"could not parse id from request": err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error retrieving event: ": err.Error()})
		return
	}

	authUserId := context.GetInt64("userId")

	if event.UserId != authUserId {
		context.JSON(http.StatusForbidden, gin.H{"message": "User not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"couldnt parse request data": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	updatedEvent.UserId = authUserId

	err = updatedEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"couldn't update event": err.Error()})
		return
	}

	// return OK response with success message.
	context.JSON(http.StatusOK, gin.H{
		"message": "event updated",
	})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("event_id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"couldn't parse id from request": err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"couldn't find event with given ID": err.Error()})
		return
	}

	authUserId := context.GetInt64("userId")

	if event.UserId != authUserId {
		context.JSON(http.StatusForbidden, gin.H{"message": "User not authorized to update event"})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"couldn't delete event": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted"})
}
