package handlers

import (
	"bytes"
	"d7024e/kademlia"
	"d7024e/models"
	"d7024e/state"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(state *state.State) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/join", func(c *gin.Context) {
		HandleJoin(c, state)
	})
	return r
}

func setupState() *state.State {
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")
	routingTable := kademlia.NewRoutingTable(contact)
	return &state.State{Node: contact, RoutingTable: routingTable}
}

// Test for missing message body
func TestJoin_MissingMessage(t *testing.T) {
	state := setupState()

	r := setupRouter(state)

	req, _ := http.NewRequest("POST", "/join", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "invalid request"}`, w.Body.String())
}

// Test for invalid message type
func TestJoin_InvalidMessageType(t *testing.T) {
	state := setupState()
	r := setupRouter(state)

	message := models.Message{
		Type: models.PING, // Invalid type
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
	state := setupState()
	r := setupRouter(state)

	sender := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8081")
	state.RoutingTable.AddContact(sender)

	message := models.Message{
		Type:    models.JOIN,
		Receiver: state.Node,
		Sender:  sender,
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
	expectedResponse := models.Message{
		Sender:   state.Node,
		Receiver: sender,
		Type:     models.ACK,
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
