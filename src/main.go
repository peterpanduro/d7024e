package main

import (
	"d7024e/handlers"
	"d7024e/kademlia"
	"d7024e/state"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	id := kademlia.NewKademliaID(kademlia.NewRandomKademliaID().String())
	state := &state.State{ID: id}
	fmt.Println("Starting node", state)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		handlers.MessageHandler(c, state)
	})
	r.GET("/ping", func(c *gin.Context) {
		handlers.Ping(c, state)
	})
	r.Run()
}
