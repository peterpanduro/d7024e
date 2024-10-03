package handlers

import (
	"d7024e/kademlia"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePing(c *gin.Context, routingTable kademlia.RoutingTable) {
	if c.Request.Method != http.MethodPost {
		response, err := kademlia.Ping(routingTable, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}
	var message kademlia.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := kademlia.Ping(routingTable, &message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
