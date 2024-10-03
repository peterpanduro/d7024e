package kademlia

import "d7024e/helpers"

type KademliaHandler interface {
	Handle(routingTable RoutingTable, message *Message) (*Message, *helpers.HTTPError)
}
