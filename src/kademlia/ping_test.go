package kademlia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingGet(t *testing.T) {
	rt := SetupRoutingTable()
	message := &Message{
		Sender:   nil,
		Receiver: nil,
		Type:     PING,
		Data:     nil,
	}
	response, _ := Ping(rt, message)
	expectedResponse := &Message{
		Sender:   rt.Me,
		Receiver: nil,
		Type:     ACK,
		Data:     nil,
	}
	assert.Equal(t, expectedResponse, response)
}

func TestPingPost(t *testing.T) {
	rt := SetupRoutingTable()
	sender := NewContact(NewRandomKademliaID(), "127.0.0.1:8081")
	message := &Message{
		Sender:   sender,
		Receiver: rt.Me,
		Type:     PING,
		Data:     nil,
	}
	response, _ := Ping(rt, message)
	expectedResponse := &Message{
		Sender:   rt.Me,
		Receiver: sender,
		Type:     ACK,
		Data:     nil,
	}
	assert.Equal(t, expectedResponse, response)
}

func TestPingInvalidType(t *testing.T) {
	rt := SetupRoutingTable()
	message := &Message{
		Sender:   rt.Me,
		Receiver: nil,
		Type:     MsgType("INVALID"),
		Data:     nil,
	}
	response, err := Ping(rt, message)
	assert.Error(t, err)
	assert.Nil(t, response)
}
