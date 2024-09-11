package main

import (
	// "d7024e/kademlia"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting node")
	// Using stuff from the kademlia package here. Something like...
	//id := kademlia.NewKademliaID("FFFFFFFF00000000000000000000000000000000")
	//contact := kademlia.NewContact(id, "localhost:8000")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
