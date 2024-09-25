package kademlia

import (
	"errors"
)

func FindNode(routingTable *RoutingTable, message *Message) (*Message, error) {
	if message == nil {
		return nil, errors.New("Message is nil")
	}
	if routingTable == nil {
		return nil, errors.New("Routing table is nil")
	}
	if message.Type != FIND_NODE {
		return nil, errors.New("Message type is not FIND_NODE")
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
