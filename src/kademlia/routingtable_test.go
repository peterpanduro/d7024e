package kademlia

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoutingTable(t *testing.T) {
	me := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	rt := NewRoutingTable(me)
	assert.Equal(t, me, rt.Me)
}

func TestRoutingTableAddFirstContact(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	contact := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8001")
	rt.AddContact(contact)
	assert.Equal(t, contact, rt.buckets[0].list.Front().Value)
}
