package handlers

import (
	"d7024e/kademlia"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFindNode(c *gin.Context, routingTable kademlia.RoutingTable) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only POST requests are allowed"})
		return
	}
	var message kademlia.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := kademlia.FindNode(routingTable, &message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

