package main

import (
	"d7024e/handlers"
	"d7024e/kademlia"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	host := initHost()
	routingTable := initNode(host)
	startServer(routingTable)
}

func initHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Couldn't get hostname")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := hostname + ":" + port
	return host
}

func initNode(host string) kademlia.RoutingTable {
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), host)
	routingTable := kademlia.NewRoutingTable(contact)
	return routingTable
}

func startServer(routingTable kademlia.RoutingTable) {
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
