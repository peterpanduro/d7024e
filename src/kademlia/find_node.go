package kademlia

import (
	"d7024e/helpers"
	"net/http"
)

type FindNodeHandler struct {}
func (findNode FindNodeHandler) Handle(routingTable RoutingTable, message *Message) (*Message, *helpers.HTTPError) {
	return FindNode(routingTable, message)
}

func FindNode(routingTable RoutingTable, message *Message) (*Message, *helpers.HTTPError) {
	if message == nil {
		return nil, helpers.NewHTTPError(http.StatusInternalServerError, "Message is nil")
	}
	if message.Type != FIND_NODE {
		return nil, helpers.NewHTTPError(http.StatusBadRequest, "Invalid message type")
	}

	sender := message.Sender
	// TODO: implement this
	// closestContacts := routingTable.FindClosestContacts(sender, 5)

	// Save the sender in the routing table
	routingTable.AddContact(sender)

	// Set the receiver to the message sender
	response := Message{
		Sender:   routingTable.Me,
		Receiver: message.Sender,
		Type:     ACK,
		Data:     nil,
	}
	return &response, nil
}
