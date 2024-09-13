package handlers

import (
	"bytes"
	"d7024e/kademlia"
	"d7024e/models"
	"d7024e/state"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessageHandler_Ping(t *testing.T) {
	// Set up the Gin engine and the router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a mock KademliaID and state for testing
	mockSenderID := kademlia.NewKademliaID(kademlia.NewRandomKademliaID().String())
	mockReceiverID := kademlia.NewKademliaID(kademlia.NewRandomKademliaID().String())
	mockedState := &state.State{
		ID: mockReceiverID,
	}

	router.POST("/", func(c *gin.Context) {
		MessageHandler(c, mockedState)
	})

	// Create a sample message that simulates a "ping" message
	message := models.Message{
		Sender:   mockSenderID.String(),
		Receiver: mockReceiverID.String(),
		Type:     models.PING,
		Data:     nil,
	}

	// Marshal the message into JSON format
	messageJSON, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Couldn't marshal message: %v", err)
	}

	// Create an HTTP POST request with the message as the body
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(messageJSON))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the HTTP request to the Gin router
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the expectedResponse to JSON
	expectedResponse := `{
		"sender": "` + mockReceiverID.String() + `",
		"receiver": "",
		"msgType": "ACK",
		"data": null
	}`

	// Check the response body
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestMessageHandler_InvalidType(t *testing.T) {
	// Set up the Gin engine and the router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a mock KademliaID and state for testing
	mockSenderID := kademlia.NewKademliaID(kademlia.NewRandomKademliaID().String())
	mockReceiverID := kademlia.NewKademliaID(kademlia.NewRandomKademliaID().String())
	appState := &state.State{
		ID: mockReceiverID,
	}

	// Replace the Ping handler with the MockPing for testing
	router.POST("/", func(c *gin.Context) {
		MessageHandler(c, appState)
	})

	// Create a sample message that simulates an invalid message type
	message := models.Message{
		Sender:   mockSenderID.String(),
		Receiver: mockReceiverID.String(),
		Type:     models.MsgType("INVALID"),
		Data:     nil,
	}

	// Marshal the message into JSON format
	messageJSON, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Couldn't marshal message: %v", err)
	}

	// Create an HTTP POST request with the message as the body
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(messageJSON))
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the HTTP request to the Gin router
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Define the expected error response
	expectedResponse := `{"error": "invalid message type"}`

	// Check the response body
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
