package handlers

import (
	"gin-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ConfirmEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("event_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "could not parse event id"})
		return
	}

	participantId, err := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "could not parse participant id"})
		return
	}

	authUserId := context.GetInt64("userId")

	if participantId != authUserId {
		context.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to sign up another user"})
		return
	}

	_, err = models.GetUsersEvent(eventId, participantId)

	if err == nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "already confirmed"})
		return
	}

	var usersEvent models.UsersEvent
	usersEvent.UserId = participantId
	usersEvent.EventId = eventId

	err = usersEvent.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store event confirmation", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user event was stored"})
}

func DeclineEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("event_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "could not parse event id"})
		return
	}

	participantId, err := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "could not parse participant id"})
		return
	}

	authUserId := context.GetInt64("userId")

	if participantId != authUserId {
		context.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to decline event for another user"})
		return
	}

	usersEvent, err := models.GetUsersEvent(eventId, participantId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "user was not already confirmed"})
		return
	}

	err = usersEvent.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete confirmation", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user event was deleted"})
}
