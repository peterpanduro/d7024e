package kademlia

import (
	"errors"
)

func Ping(routingTable *RoutingTable, message *Message) (*Message, error) {
	if message == nil {
		// GET method. No reciever.
		response := &Message{
			Sender:   routingTable.Me,
			Receiver: nil,
			Type:     ACK,
			Data:     nil,
		}
		return response, nil
	}

	if message.Type != PING {
		return nil, errors.New("Invalid message type")
	}

	response := &Message{
		Sender:   routingTable.Me,
		Receiver: message.Sender,
		Type:     ACK,
		Data:     nil,
	}
	return response, nil
}
