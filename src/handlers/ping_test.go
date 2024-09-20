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

func TestPingGet(t *testing.T) {
	// Set up the Gin engine and the router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mock the state
	rt := setupRoutingTable()

	// Wrap the Ping handler with the appState
	router.GET("/ping", func(c *gin.Context) {
		HandlePing(c, rt)
	})

	// Create an HTTP request to test the Ping handler
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Serve the HTTP request to the Gin router
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Define the expected response as a Message struct
	expectedResponse := kademlia.Message{
		Sender:   rt.Me,
		Receiver: nil,
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

func TestPingPost(t *testing.T) {
	// Set up the Gin engine and the router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mock the state
	rt := setupRoutingTable()
	mockSender := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8081")

	// Wrap the Ping handler with the appState
	router.POST("/ping", func(c *gin.Context) {
		HandlePing(c, rt)
	})

	message := kademlia.Message{
		Sender:   mockSender,
		Receiver: rt.Me,
		Type:     kademlia.PING,
		Data:     nil,
	}
	messageJSON, err := json.Marshal(message)

	// Create an HTTP request to test the Ping handler
	req, err := http.NewRequest(http.MethodPost, "/ping", bytes.NewBuffer(messageJSON))
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

	expectedResponse := kademlia.Message{
		// Define the expected response as a Message struct
		Sender:   rt.Me,
		Receiver: mockSender,
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
