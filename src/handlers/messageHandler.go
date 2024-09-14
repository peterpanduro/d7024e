package handlers

import (
	"d7024e/models"
	"d7024e/state"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageHandler(c *gin.Context, state *state.State) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message format"})
		return
	}

	switch message.Type {
	case models.PING:
		Ping(c, state)
	default:
		c.JSON(http.StatusOK, gin.H{"error": "invalid message type"})
	}
}
