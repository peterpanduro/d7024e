/*
	package main

import (

	"d7024e/handlers"
	"d7024e/kademlia"
	"log"
	"os"

	"github.com/gin-gonic/gin"
<<<<<<< HEAD
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
=======

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
*/
package main

import (
	"d7024e/kademlia"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Create a contact for the Kademlia network
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")

	// Create a new instance of Kademlia
	kad := kademlia.NewKademlia(contact)

	// Handle CLI arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: put <file_path>")
		return
	}

	command := os.Args[1]

	switch command {
	case "put":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a file path to upload")
			return
		}

		filePath := os.Args[2]
		// Read the file content
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading the file:", err)
			return
		}

		// Store the content using Kademlia and get the hash
		hash := kad.Store(content)

		// Output the hash of the object
		fmt.Printf("File uploaded successfully. Hash: %s\n", hash)

	default:
		fmt.Println("Invalid command")
	}
>>>>>>> f1ede7a (Update main.go)
}
