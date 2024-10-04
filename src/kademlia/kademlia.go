package kademlia

import (
	"crypto/sha256"
	"d7024e/helpers"
	"encoding/hex"
	"time"
)

// KademliaHandler interface defines the method signature required for handling messages.
type KademliaHandler interface {
	Handle(routingTable RoutingTable, message *Message) (*Message, *helpers.HTTPError)
}

// Kademlia represents the main Kademlia struct.
type Kademlia struct {
	routingTable *RoutingTable
	storage      map[string][]byte // Simple in-memory storage for data
}

// NewKademlia returns a new Kademlia instance.
func NewKademlia(contact *Contact) *Kademlia {
	return &Kademlia{
		routingTable: NewRoutingTable(contact),
		storage:      make(map[string][]byte),
	}
}

// Handle processes incoming messages based on the message type.
func (kademlia *Kademlia) Handle(routingTable RoutingTable, message *Message) (*Message, *helpers.HTTPError) {
	switch message.Type {
	case FIND_NODE:
		// Lookup the closest contacts for the given contact
		closestContacts := kademlia.LookupContact(message.Sender)
		return &Message{Contacts: closestContacts}, nil

	case FIND_VALUE:
		// Lookup data based on the provided hash
		data := kademlia.LookupData(*message.Data.HASH)
		if data == nil {
			return nil, &helpers.HTTPError{Message: "Data not found"}
		}
		return &Message{Data: data}, nil

	case STORE:
		// Store the data locally
		kademlia.Store(*message.Data.VALUE)
		return &Message{Ack: true}, nil

	default:
		// Handle unknown message types
		return nil, &helpers.HTTPError{Message: "Unknown message type"}
	}
}

// LookupContact looks for the closest contacts to the target in the Kademlia network.
func (kademlia *Kademlia) LookupContact(target *Contact) []*Contact {
	// Finds the closest contacts to the target node using the routing table.
	closestContacts := kademlia.routingTable.FindClosestContacts(target, bucketSize)
	return closestContacts
}

// LookupData attempts to find data associated with the provided hash in the network.
func (kademlia *Kademlia) LookupData(hash string) *MsgData {
	// Check if the data is stored locally first
	if value, exists := kademlia.storage[hash]; exists {
		return &MsgData{HASH: &hash, VALUE: &value}
	}

	// Create a message to request the data associated with the hash
	message := Message{
		Sender: kademlia.routingTable.Me,
		Type:   FIND_VALUE,
		Data:   &MsgData{HASH: &hash},
	}

	// Get the closest contacts
	closestContacts := kademlia.routingTable.FindClosestContacts(kademlia.routingTable.Me, bucketSize)

	// Channel to collect responses
	responseChan := make(chan *MsgData, len(closestContacts))

	// Send requests concurrently to multiple nodes
	for _, contact := range closestContacts {
		go func(contact *Contact) {
			message.Receiver = contact
			err := SendMessage(message)
			if err != nil {
				// Handle error, ignoring it in this case
				return
			}

			// Simulate receiving data from the contact (placeholder logic)
			mockData := []byte("mock data")
			receivedData := &MsgData{HASH: &hash, VALUE: &mockData}
			responseChan <- receivedData
		}(contact)
	}

	// Wait for one response or timeout
	select {
	case response := <-responseChan:
		return response
	case <-time.After(5 * time.Second): // Timeout after 5 seconds
		return nil
	}
}

// Store adds data to the local storage and broadcasts the store message to the closest contacts.
func (kademlia *Kademlia) Store(data []byte) {
	// Generate the hash for the data to be stored
	hash := GenerateHash(data)

	// Store data locally
	kademlia.storage[hash] = data

	// Create a message to store the data
	message := Message{
		Sender: kademlia.routingTable.Me,
		Type:   STORE,
		Data:   &MsgData{HASH: &hash, VALUE: &data},
	}

	// Get the closest contacts to the data's hash
	closestContacts := kademlia.routingTable.FindClosestContacts(kademlia.routingTable.Me, bucketSize)

	// Send store message to each contact concurrently
	for _, contact := range closestContacts {
		go func(contact *Contact) {
			message.Receiver = contact
			err := SendMessage(message)
			if err != nil {
				// Handle network errors or unsuccessful sends but we just ignore it in this case
				return
			}

			// Simulate acknowledgment handling (placeholder for now)
		}(contact)
	}
}

// GenerateHash generates the SHA-256 hash of the provided data.
func GenerateHash(data []byte) string {
	// Create a new SHA-256 hash object
	hash := sha256.New()

	// Write the data to the hash
	hash.Write(data)

	// Calculate the hash
	hashSum := hash.Sum(nil)

	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(hashSum)
}
