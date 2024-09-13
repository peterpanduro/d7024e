package handlers

import (
	"d7024e/kademlia"
	"d7024e/models" // Import the models package where Message is defined
	"d7024e/state"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// Set up the Gin engine and the router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a mock KademliaID for testing
	mockID := kademlia.NewKademliaID(kademlia.NewRandomKademliaID().String())
	appState := &state.State{
		ID: mockID,
	}

	// Wrap the Ping handler with the appState
	router.GET("/ping", func(c *gin.Context) {
		Ping(c, appState) // Pass the appState to the handler
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
	expectedResponse := models.Message{
		Sender:   mockID.String(),
		Receiver: "",
		Type:     models.PONG,
		Data:     nil,
	}

	// Convert the expectedResponse to JSON
	expectedJSON := `{
		"sender": "` + expectedResponse.Sender + `",
		"receiver": "",
		"msgType": "PONG",
		"data": null
	}`

	// Check the response body
	assert.JSONEq(t, expectedJSON, w.Body.String())
}
