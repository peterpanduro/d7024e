package main

import (
	"d7024e/handlers"
	"d7024e/kademlia"
	"d7024e/state"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")
	state := &state.State{Node: &contact}
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
	r.Run()
}
