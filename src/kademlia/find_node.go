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
	FindNodeHash := message.Data.HASH
	if FindNodeHash == nil {
		return nil, helpers.NewHTTPError(http.StatusBadRequest, "No HASH Data in FIND_NODE message")
	}

	sender := message.Sender
	routingTable.AddContact(sender)

	virtualContact := &Contact{ID: NewKademliaID(*FindNodeHash)}
	closestContacts := routingTable.FindClosestContacts(virtualContact, 1)
	closestContact := closestContacts[0]
	ClosestContactHash := closestContact.ID.String()

	// Set the receiver to the message sender
	response := Message{
		Sender:   routingTable.Me,
		Receiver: message.Sender,
		Type:     FOUND_NODE,
		Data:     &MsgData{HASH: &ClosestContactHash, VALUE: nil, ERR: nil},
	}
	return &response, nil
}
