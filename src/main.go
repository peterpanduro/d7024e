package main

import (
	"d7024e/handlers"
	"d7024e/kademlia"
	"d7024e/state"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")
	routingTable := kademlia.NewRoutingTable(contact)
	state := &state.State{Node: contact, RoutingTable: routingTable}

	// TEMPORARY
	for i := 0; i < 2; i++ {
		contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:"+strconv.Itoa(8080+i))
		state.RoutingTable.AddContact(contact)
	}

	fmt.Println("Starting node", state)
	r := gin.Default()
	r.POST("", func(c *gin.Context) {
		handlers.MessageHandler(c, state)
	})
	r.GET("/ping", func(c *gin.Context) {
		handlers.HandlePing(c, state)
	})
	r.POST("/ping", func(c *gin.Context) {
		handlers.HandlePing(c, state)
	})
	r.POST("/join", func(c *gin.Context) {
		handlers.HandleJoin(c, state)
	})
	r.Run()
}
