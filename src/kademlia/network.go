package kademlia

import (
	"encoding/json"
	"net"
)

// Serialize the message to JSON
func SendMessage(message Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Resolve the receiver's address
	addr, err := net.ResolveUDPAddr("udp", message.Receiver.Address)
	if err != nil {
		return err
	}

	// Create a UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send the message
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}
