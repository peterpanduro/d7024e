package handlers

import (
	"d7024e/kademlia"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageHandler(c *gin.Context, routingTable kademlia.RoutingTable) {
	var message kademlia.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid message format",
		})
		return
	}

	var handler kademlia.KademliaHandler
	switch message.Type {
	case kademlia.PING:
		handler = kademlia.PingHandler{}
	case kademlia.FIND_NODE:
		handler = kademlia.FindNodeHandler{}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid message type"})
		return
	}

	response, err := handler.Handle(routingTable, &message)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
