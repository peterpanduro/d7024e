package main

import (
	"d7024e/handlers"
	"d7024e/kademlia"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")
	routingTable := kademlia.NewRoutingTable(contact)

	fmt.Println("Starting node", routingTable)
	r := gin.Default()
	r.POST("", func(c *gin.Context) {
		handlers.MessageHandler(c, routingTable)
	})
	r.GET("/ping", func(c *gin.Context) {
		handlers.HandlePing(c, routingTable)
	})
	r.POST("/ping", func(c *gin.Context) {
		handlers.HandlePing(c, routingTable)
	})
	r.Run()
}
