package kademlia

func SetupRoutingTable() RoutingTable {
	contact := NewContact(NewKademliaID("1122334455667788990011223344556677889900"), "127.0.0.1:8080")
	routingTable := NewRoutingTable(contact)
	return routingTable
}

