package handlers

import (
	"d7024e/kademlia"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePing(c *gin.Context, routingTable *kademlia.RoutingTable) {
	if c.Request.Method == http.MethodPost {
		var message kademlia.Message
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Ping(c, routingTable, &message)
	} else {
		Ping(c, routingTable, nil)
	}
}

func Ping(c *gin.Context, routingTable *kademlia.RoutingTable, message *kademlia.Message) {
	// POST request without message body
	if message == nil && c.Request.Method != http.MethodGet {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing message body"})
		return
	}

	// GET request
	if message == nil {
		response := kademlia.Message{
			Sender:   routingTable.Me,
			Receiver: nil,
			Type:     kademlia.ACK,
			Data:     nil,
		}
		c.JSON(http.StatusOK, response)
		return
	}

	// Set the receiver to the message sender
	response := kademlia.Message{
		Sender:   routingTable.Me,
		Receiver: message.Sender,
		Type:     kademlia.ACK,
		Data:     nil,
	}
	c.JSON(http.StatusOK, response)
}
