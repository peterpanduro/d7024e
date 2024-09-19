package handlers

import (
	"d7024e/kademlia"
	"d7024e/models"
	"d7024e/state"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleJoin(c *gin.Context, state *state.State) {
	if c.Request.Method == http.MethodPost {
		var message models.Message
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Join(c, state, &message)
	} else {
		Join(c, state, nil)
	}
}

func Join(c *gin.Context, state *state.State, message *models.Message) {
	// POST request without message body
	if message == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing message body"})
		return
	}

	if message.Type != models.JOIN {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message type"})
		return
	}

	// Find the closest contacts to the sender
	sender := message.Sender
	println(sender)
	closestContacts := state.RoutingTable.FindClosestContacts(sender, 5)
	println(closestContacts)

	// Save the sender in the routing table
	// state.RoutingTable.AddContact(sender)

	// Set the receiver to the message sender
	response := models.Message{
		Sender:   state.Node,
		Receiver: message.Sender,
		Type:     models.ACK,
		Data:     nil,
	}
	c.JSON(http.StatusOK, response)

	// Send FIND_NODE messages to the sender
	// for _, contact := range closestContacts {
	// 	//
	// 	SendFindContactMessage(&contact, state.Node, sender)
	// }
}

func SendFindContactMessage(contact *kademlia.Contact, sender *kademlia.Contact, receiver *kademlia.Contact) {
	bytes := contact.ID[:]
	data := models.MsgData{
		VALUE: &bytes,
	}
	message := models.Message{
		Sender:   sender,
		Receiver: receiver,
		Type:     models.FIND_NODE,
		Data:     &data,
	}
	fmt.Println(message)
}
