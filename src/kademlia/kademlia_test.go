package kademlia

import (
	"strconv"
	"testing"
)

// TestNewKademlia tests the creation of a new Kademlia instance.
func TestNewKademlia(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "127.0.0.1:8080")
	kademlia := NewKademlia(contact)

	if kademlia == nil || kademlia.routingTable == nil {
		t.Errorf("NewKademlia() did not initialize properly")
	}
}

// TestLookupContact tests the LookupContact function.
func TestLookupContact(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "127.0.0.1:8080")
	kademlia := NewKademlia(contact)

	// Add some mock contacts to the routing table
	for i := 0; i < 3; i++ {
		mockContact := NewContact(NewRandomKademliaID(), "127.0.0.1:808"+strconv.Itoa(i))
		kademlia.routingTable.AddContact(mockContact)
	}

	// Lookup a random contact
	target := NewContact(NewRandomKademliaID(), "127.0.0.1:8090")
	closestContacts := kademlia.LookupContact(target)

	if len(closestContacts) == 0 {
		t.Errorf("LookupContact() did not find any closest contacts")
	}
}

// TestLookupData tests the LookupData function.
func TestLookupData(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "127.0.0.1:8080")
	kademlia := NewKademlia(contact)

	// Store some data locally to simulate existing data
	data := []byte("sample data")
	hash := GenerateHash(data)
	kademlia.storage[hash] = data

	// Test LookupData for locally stored data
	msgData := kademlia.LookupData(hash)
	if msgData == nil || msgData.VALUE == nil || string(*msgData.VALUE) != string(data) {
		t.Errorf("LookupData() failed to find locally stored data")
	}

	// Test LookupData for non-existent data (expecting nil response)
	msgData = kademlia.LookupData("non_existent_hash")
	if msgData != nil {
		t.Errorf("LookupData() should return nil for non-existent data")
	}
}

// TestStore tests the Store function.
func TestStore(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "127.0.0.1:8080")
	kademlia := NewKademlia(contact)

	// Mock data to be stored
	data := []byte("sample data to store")

	// Store the data in the network
	kademlia.Store(data)

	// Verify that the data is stored locally
	hash := GenerateHash(data)
	if storedData, exists := kademlia.storage[hash]; !exists || string(storedData) != string(data) {
		t.Errorf("Store() failed to store data locally")
	}
}

// TestGenerateHash tests the GenerateHash function.
func TestGenerateHash(t *testing.T) {
	data := []byte("some test data")
	hash := GenerateHash(data)

	// We expect a non-empty string as the hash
	if hash == "" {
		t.Errorf("GenerateHash() returned an empty string")
	}

	// Correct expected SHA-256 hash of "some test data"
	expectedHash := "f70c5e847d0ea29088216d81d628df4b4f68f3ccabb2e4031c09cc4d129ae216"
	if hash != expectedHash {
		t.Errorf("GenerateHash() returned wrong hash, got %s, expected %s", hash, expectedHash)
	}
}
