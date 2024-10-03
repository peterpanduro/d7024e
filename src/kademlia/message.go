package kademlia

type Message struct {
	Sender   *Contact `json:"sender"`
	Receiver *Contact `json:"receiver"`
	Type     MsgType           `json:"msgType"`
	Data     *MsgData          `json:"data"`
}

type MsgData struct {
	HASH  *string `json:"hash"`
	VALUE *[]byte `json:"value"`
	ERR   *string `json:"err"`
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
