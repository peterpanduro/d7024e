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

func TestFindNode_MissingMessage(t *testing.T) {
	rt := kademlia.SetupRoutingTable()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/find_node", func(c *gin.Context) {
		HandleFindNode(c, rt)
	})

	req, _ := http.NewRequest("POST", "/find_node", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "invalid request"}`, w.Body.String())
}

func TestFindNode_InvalidMessageType(t *testing.T) {
	rt := kademlia.SetupRoutingTable()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/find_node", func(c *gin.Context) {
		HandleFindNode(c, rt)
	})

	message := kademlia.Message{
		Type: kademlia.PING, // Invalid type
	}
	jsonValue, _ := json.Marshal(message)

	req, _ := http.NewRequest("POST", "/find_node", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "Invalid message type"}`, w.Body.String())
}
