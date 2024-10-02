package handlers

import (
	"bytes"
	"d7024e/kademlia"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMessageHandler_Ping(t *testing.T) {
	// Set up the Gin engine and the router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a mock KademliaID and state for testing
	rt := kademlia.SetupRoutingTable()
	sender := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8081")

	router.POST("/", func(c *gin.Context) {
		MessageHandler(c, rt)
	})

	message := kademlia.Message{
		Sender:   sender,
		Receiver: rt.Me,
		Type:     kademlia.PING,
		Data:     nil,
	}
	messageJSON, err := json.Marshal(message)

	// Create an HTTP request to test the Message handler
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

func TestMessageHandler_InvalidType(t *testing.T) {
	// Set up the Gin engine and the router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a mock KademliaID and state for testing
	rt := kademlia.SetupRoutingTable()
	sender := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8081")

	// Replace the Ping handler with the MockPing for testing
	router.POST("/", func(c *gin.Context) {
		MessageHandler(c, rt)
	})

	// Create a sample message that simulates an invalid message type
	message := kademlia.Message{
		Sender:   rt.Me,
		Type:     kademlia.MsgType("INVALID"),
		Receiver: sender,
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
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Define the expected error response
	expectedResponse := `{"error": "invalid message type"}`

	// Check the response body
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
