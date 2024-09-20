package handlers

import (
	"bytes"
	"d7024e/kademlia"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func loadMockData(routingTable *kademlia.RoutingTable) {
	for i := 0; i < 5; i++ {
		contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:"+strconv.Itoa(8080+i))
		routingTable.AddContact(contact)
	}
}

func setupRouter(routingTable *kademlia.RoutingTable) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/join", func(c *gin.Context) {
		HandleJoin(c, routingTable)
	})
	return r
}

func setupRoutingTable() *kademlia.RoutingTable {
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")
	return kademlia.NewRoutingTable(contact)
}

// Test for missing message body
func TestJoin_MissingMessage(t *testing.T) {
	rt := setupRoutingTable()
	r := setupRouter(rt)

	req, _ := http.NewRequest("POST", "/join", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "invalid request"}`, w.Body.String())
}

// Test for invalid message type
func TestJoin_InvalidMessageType(t *testing.T) {
	rt := setupRoutingTable()
	r := setupRouter(rt)

	message := kademlia.Message{
		Type: kademlia.PING, // Invalid type
	}
	jsonValue, _ := json.Marshal(message)

	req, _ := http.NewRequest("POST", "/join", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "Invalid message type"}`, w.Body.String())
}

// Test for valid join request
func TestJoin_ValidMessageReturnsACK(t *testing.T) {
	rt := setupRoutingTable()
	r := setupRouter(rt)

	sender := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8081")
	message := kademlia.Message{
		Type:     kademlia.JOIN,
		Receiver: rt.Me,
		Sender:   sender,
	}
	jsonValue, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Couldn't marshal message: %v", err)
		return
	}

	req, err := http.NewRequest("POST", "/join", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Define the expected response as a Message struct
	expectedResponse := kademlia.Message{
		Sender:   rt.Me,
		Receiver: sender,
		Type:     kademlia.ACK,
		Data:     nil,
	}

	// Convert the expectedResponse to JSON
	expectedResponseJSON, err := json.Marshal(expectedResponse)
	if err != nil {
		t.Fatalf("Couldn't marshal expected response: %v", err)
	}

	// Check the response body
	assert.JSONEq(t, string(expectedResponseJSON), w.Body.String())
}

func TestJoin_JoinEndsUpInRoutingTable(t *testing.T) {
	// TODO: Write tests after bucket is properly tested
}
