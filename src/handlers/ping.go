package handlers

import (
	"d7024e/models"
	"d7024e/state"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlePing(c *gin.Context, state *state.State) {
	if c.Request.Method == http.MethodPost {
		var message models.Message
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Ping(c, state, &message)
	} else {
		Ping(c, state, nil)
	}
}

func Ping(c *gin.Context, state *state.State, message *models.Message) {
	// POST request without message body
	if message == nil && c.Request.Method != http.MethodGet {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing message body"})
		return
	}

	// GET request
	if message == nil {
		response := models.Message{
			Sender:   state.Node,
			Receiver: nil,
			Type:     models.ACK,
			Data:     nil,
		}
		c.JSON(http.StatusOK, response)
		return
	}

	// Set the receiver to the message sender
	response := models.Message{
		Sender:   state.Node,
		Receiver: message.Sender,
		Type:     models.ACK,
		Data:     nil,
	}
	c.JSON(http.StatusOK, response)
}
