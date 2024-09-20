package handlers

import (
	"d7024e/kademlia"
	"net/http"
	"github.com/gin-gonic/gin"
)

func MessageHandler(c *gin.Context, routingTable *kademlia.RoutingTable) {
	var message kademlia.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid message format",
		})
		return
	}

	switch message.Type {
	case kademlia.PING:
		Ping(c, routingTable, &message)
	case kademlia.JOIN:
		Join(c, routingTable, &message)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid message type"})
	}
}
