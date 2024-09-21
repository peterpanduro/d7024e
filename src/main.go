package main

import (
	"d7024e/handlers"
	"d7024e/kademlia"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoadMockData(routingTable *kademlia.RoutingTable) {
	// TEMPORARY
	for i := 0; i < 2; i++ {
		contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:"+strconv.Itoa(8080+i))
		routingTable.AddContact(contact)
	}
}

func main() {
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")
	routingTable := kademlia.NewRoutingTable(contact)

	LoadMockData(routingTable)

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
	r.POST("/join", func(c *gin.Context) {
		handlers.HandleJoin(c, routingTable)
	})
	r.Run()
}
