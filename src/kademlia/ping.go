package kademlia

import (
	"d7024e/helpers"
	"net/http"
)

type PingHandler struct {}
func (ping PingHandler) Handle(routingTable RoutingTable, message *Message) (*Message, *helpers.HTTPError) {
	return Ping(routingTable, message)
}

func Ping(routingTable RoutingTable, message *Message) (*Message, *helpers.HTTPError) {
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
		return nil, helpers.NewHTTPError(http.StatusBadRequest, "Invalid message type")
	}

	response := &Message{
		Sender:   routingTable.Me,
		Receiver: message.Sender,
		Type:     ACK,
		Data:     nil,
	}
	return response, nil
}
