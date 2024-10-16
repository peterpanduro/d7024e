/*
	package main

import (

	"d7024e/handlers"
	"d7024e/kademlia"
	"log"
	"os"

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
*/
package main

import (
	"bufio"
	"d7024e/kademlia"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Create a contact for the Kademlia network
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")

	// Create a new instance of Kademlia
	kad := kademlia.NewKademlia(contact)

	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt the user for a command
		fmt.Print("> ")
		commandLine, _ := reader.ReadString('\n')
		handleCommand(commandLine, kad)
	}
}

func handleCommand(commandLine string, kad *kademlia.Kademlia) {
	commandLine = strings.TrimSpace(commandLine)
	args := strings.Split(commandLine, " ")

	if len(args) < 1 {
		return
	}

	command := args[0]

	switch command {
	case "put":
		if len(args) < 2 {
			fmt.Println("Please provide a file path to upload")
			return
		}

		filePath := args[1]
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading the file:", err)
			return
		}

		// Store the content using Kademlia and get the hash
		hash := kad.Store(content)
		fmt.Printf("File uploaded successfully. Hash: %s\n", hash)

	case "get":
		if len(args) < 2 {
			fmt.Println("Please provide a hash to retrieve")
			return
		}

		hash := args[1]
		// Lookup the data using the Kademlia network
		data := kad.LookupData(hash)

		if data == nil {
			fmt.Println("Data not found for the given hash")
			return
		}

		// Output the retrieved data
		fmt.Printf("Data retrieved: %s\n", string(*data.VALUE))
		fmt.Printf("Retrieved from hash: %s\n", *data.HASH)

	case "exit":
		fmt.Println("Terminating the node.")
		os.Exit(0) // Gracefully exit the program

	default:
		fmt.Println("Invalid command")
	}
}
