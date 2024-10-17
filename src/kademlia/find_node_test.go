package kademlia

import (
	"net/http"
	"testing"
)

func TestFindNode_FindSelf(t *testing.T) {
	// Setup
	senderId := "11223344556677889900112233445566778899aa"
	meId := "ff22334455667788990011223344556677889900"
	sender := &Contact{ID: NewKademliaID(senderId)}
	me := &Contact{ID: NewKademliaID(meId)}
	routingTable := NewRoutingTable(me)
	routingTable.AddContact(sender)

	// Create a valid FIND_NODE message
	message := &Message{
		Sender: sender,
		Type:   FIND_NODE,
		Data:   &MsgData{HASH: &meId, VALUE: nil, ERR: nil},
	}

	// Call the function
	response, err := FindNode(routingTable, message)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	if response == nil {
		t.Fatalf("Expected a response, but got nil")
	}
	if response.Type != FOUND_NODE {
		t.Errorf("Expected response type %v, but got %v", FOUND_NODE, response.Type)
	}
	if response.Sender.ID.String() != meId {
		t.Errorf("Expected sender to be %v, but got %v", meId, response.Sender.ID.String())
	}
	if response.Receiver.ID.String() != senderId {
		t.Errorf("Expected receiver to be %v, but got %v", senderId, response.Receiver.ID.String())
	}
	if *response.Data.HASH != senderId {
		t.Errorf("Expected closest contact ID %v, but got %v", senderId, *response.Data.HASH)
	}
}

func TestFindNode_FindClosest(t *testing.T) {
	senderId := "11223344556677889900112233445566778899aa"
	closestId := "11223344556677889900112233445566778899ab"
	meId := "ff22334455667788990011223344556677889900"
	sender := &Contact{ID: NewKademliaID(senderId)}
	me := &Contact{ID: NewKademliaID(meId)}
	closest := &Contact{ID: NewKademliaID(closestId)}
	routingTable := NewRoutingTable(me)
	routingTable.AddContact(sender)
	routingTable.AddContact(closest)

	// Create a valid FIND_NODE message
	message := &Message{
		Sender: sender,
		Type:   FIND_NODE,
		Data:   &MsgData{HASH: &closestId, VALUE: nil, ERR: nil},
	}

	// Call the function
	response, err := FindNode(routingTable, message)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	if response == nil {
		t.Fatalf("Expected a response, but got nil")
	}
	if response.Type != FOUND_NODE {
		t.Errorf("Expected response type %v, but got %v", FOUND_NODE, response.Type)
	}
	if response.Sender.ID.String() != meId {
		t.Errorf("Expected sender to be %v, but got %v", meId, response.Sender.ID.String())
	}
	if response.Receiver.ID.String() != senderId {
		t.Errorf("Expected receiver to be %v, but got %v", senderId, response.Receiver.ID.String())
	}
	if *response.Data.HASH != closestId {
		t.Errorf("Expected closest contact ID %v, but got %v", closestId, *response.Data.HASH)
	}
}
func TestFindNode_InvalidMessageType(t *testing.T) {
	// Setup
	senderId := "11223344556677889900112233445566778899aa"
	meId := "ff22334455667788990011223344556677889900"
	sender := &Contact{ID: NewKademliaID(senderId)}
	me := &Contact{ID: NewKademliaID(meId)}
	routingTable := NewRoutingTable(me)
	routingTable.AddContact(sender)

	// Create a message with an invalid type
	message := &Message{
		Sender: sender,
		Type:   STORE,
	}

	// Call the function
	_, err := FindNode(routingTable, message)

	// Assertions
	if err == nil {
		t.Fatalf("Expected an error, but got nil")
	}
	if err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %v, but got %v", http.StatusBadRequest, err.Code)
	}
	if err.Message != "Invalid message type" {
		t.Errorf("Expected error message 'Invalid message type', but got %v", err.Message)
	}
}

func TestFindNode_NilMessage(t *testing.T) {
	// Setup
	senderId := "11223344556677889900112233445566778899aa"
	meId := "ff22334455667788990011223344556677889900"
	sender := &Contact{ID: NewKademliaID(senderId)}
	me := &Contact{ID: NewKademliaID(meId)}
	routingTable := NewRoutingTable(me)
	routingTable.AddContact(sender)

	// Call the function with nil message
	_, err := FindNode(routingTable, nil)

	// Assertions
	if err == nil {
		t.Fatalf("Expected an error, but got nil")
	}
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %v, but got %v", http.StatusInternalServerError, err.Code)
	}
	if err.Message != "Message is nil" {
		t.Errorf("Expected error message 'Message is nil', but got %v", err.Message)
	}
}

