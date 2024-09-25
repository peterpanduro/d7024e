package handlers

import (
	"d7024e/kademlia"
	"github.com/gin-gonic/gin"
	"net/http"
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
		response, err := kademlia.Ping(routingTable, &message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
		return
	case kademlia.FIND_NODE:
		response, err := kademlia.FindNode(routingTable, &message)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid message type"})
	}
}
