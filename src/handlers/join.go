package handlers

import (
	"d7024e/kademlia"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleJoin(c *gin.Context, routingTable *kademlia.RoutingTable) {
	if c.Request.Method == http.MethodPost {
		var message kademlia.Message
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Join(c, routingTable, &message)
	} else {
		Join(c, routingTable, nil)
	}
}

func Join(c *gin.Context, routingTable *kademlia.RoutingTable, message *kademlia.Message) {
	// POST request without message body
	if message == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing message body"})
		return
	}

	if message.Type != kademlia.JOIN {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message type"})
		return
	}

	// Find the closest contacts to the sender
	sender := message.Sender
	closestContacts := routingTable.FindClosestContacts(sender, 5)

	// Save the sender in the routing table
	routingTable.AddContact(sender)

	// Set the receiver to the message sender
	response := kademlia.Message{
		Sender:   routingTable.Me,
		Receiver: message.Sender,
		Type:     kademlia.ACK,
		Data:     nil,
	}
	c.JSON(http.StatusOK, response)

	// Send FIND_NODE messages to the sender
	for _, contact := range closestContacts {
		SendFindContactMessage(contact, routingTable.Me, sender)
	}
}

func SendFindContactMessage(contact *kademlia.Contact, sender *kademlia.Contact, receiver *kademlia.Contact) {
	bytes := contact.ID[:]
	data := kademlia.MsgData{
		VALUE: &bytes,
	}
	message := kademlia.Message{
		Sender:   sender,
		Receiver: receiver,
		Type:     kademlia.FIND_NODE,
		Data:     &data,
	}
	kademlia.SendMessage(message)
}
