package kademlia

import (
	"testing"
	"net"
	"encoding/json"
	"time"
)

func TestSendMessage(t *testing.T) {
	// Create a mock receiver contact
	receiver := &Contact{
		ID:      NewKademliaID("ffffffff00000000000000000000000000000000"),
		Address: "localhost:8080", // Test UDP address
	}

	// Create a mock sender contact
	sender := &Contact{
		ID:      NewKademliaID("1111111100000000000000000000000000000000"),
		Address: "localhost:8081",
	}

	// Create a message
	message := Message{
		Sender:   sender,
		Receiver: receiver,
		Type:     PING,
		Data: &MsgData{
			HASH:  nil,
			VALUE: nil,
			ERR:   nil,
		},
	}

	// Set up a UDP listener to mock the receiver
	addr, err := net.ResolveUDPAddr("udp", receiver.Address)
	if err != nil {
		t.Fatalf("failed to resolve UDP address: %v", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("failed to listen on UDP address: %v", err)
	}
	defer conn.Close()

	// Create a channel to handle message reception
	done := make(chan bool)

	// Start a goroutine to receive the message
	go func() {
		// Buffer for incoming data
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			t.Errorf("error reading UDP message: %v", err)
			done <- false
			return
		}

		// Unmarshal received data into a Message object
		var receivedMessage Message
		if err := json.Unmarshal(buffer[:n], &receivedMessage); err != nil {
			t.Errorf("failed to unmarshal message: %v", err)
			done <- false
			return
		}

		// Assert that the received message matches the sent one
		if receivedMessage.Type != message.Type {
			t.Errorf("received incorrect message type: got %+v, want %+v", receivedMessage.Type, message.Type)
			done <- false
			return
		}

		if *receivedMessage.Sender.ID != *message.Sender.ID {
			t.Errorf("received incorrect message sender: got %+v, want %+v", receivedMessage.Sender.ID, message.Sender.ID)
			done <- false
			return
		}

		// Compare Data field values instead of pointers
		if (receivedMessage.Data == nil && message.Data != nil) || (receivedMessage.Data != nil && message.Data == nil) {
			t.Errorf("received message data is nil or not nil, mismatch: got %+v, want %+v", receivedMessage.Data, message.Data)
			done <- false
			return
		}

		// Check HASH, VALUE, and ERR fields
		if receivedMessage.Data != nil {
			if receivedMessage.Data.HASH != nil && message.Data.HASH != nil && *receivedMessage.Data.HASH != *message.Data.HASH {
				t.Errorf("HASH mismatch: got %v, want %v", *receivedMessage.Data.HASH, *message.Data.HASH)
			}
			if receivedMessage.Data.ERR != nil && message.Data.ERR != nil && *receivedMessage.Data.ERR != *message.Data.ERR {
				t.Errorf("ERR mismatch: got %v, want %v", *receivedMessage.Data.ERR, *message.Data.ERR)
			}
			if receivedMessage.Data.VALUE != nil && message.Data.VALUE != nil && string(*receivedMessage.Data.VALUE) != string(*message.Data.VALUE) {
				t.Errorf("VALUE mismatch: got %v, want %v", string(*receivedMessage.Data.VALUE), string(*message.Data.VALUE))
			}
		}

		done <- true
	}()

	// Send the message using the SendMessage function
	err = SendMessage(message)
	if err != nil {
		t.Fatalf("SendMessage failed: %v", err)
	}

	// Wait for the receiver goroutine to finish
	select {
	case success := <-done:
		if !success {
			t.Fatal("message reception failed")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("test timed out waiting for message reception")
	}
}
