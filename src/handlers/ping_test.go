package handlers

import (
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
	router.GET("/ping", Ping)

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

	// Check the response body
	expectedResponse := `"PONG"`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
