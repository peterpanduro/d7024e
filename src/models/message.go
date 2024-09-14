package models

import "d7024e/kademlia"

type Message struct {
	Sender   *kademlia.Contact `json:"sender"`
	Receiver *kademlia.Contact `json:"receiver"`
	Type     MsgType           `json:"msgType"`
	Data     *MsgData          `json:"data"`
}

type MsgData struct {
	STORE []byte `json:"store"`
	HASH  string `json:"hash"`
	// NODE  KademliaID        `json:"node"`
	// NODES ContactCandidates `json:"nodes"`
	VALUE []byte `json:"value"`
	ERR   string `json:"err"`
}

type MsgType string

const (
	PING        MsgType = "PING"
	ACK         MsgType = "ACK"
	STORE       MsgType = "STORE"
	STORED      MsgType = "STORED"
	FIND_NODE   MsgType = "FIND_NODE"
	FOUND_NODE  MsgType = "FOUND_NODE"
	FIND_VALUE  MsgType = "FIND_VALUE"
	FOUND_VALUE MsgType = "FOUND_VALUE"
	ERR         MsgType = "ERR"
	GET         MsgType = "GET"
)
