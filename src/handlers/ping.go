package handlers

import (
	"d7024e/models"
	"net/http"
	"d7024e/state"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context, state *state.State) {
	response := models.Message{
		Sender:   state.ID.String(),
		Receiver: c.Param("sender"),
		Type:     models.ACK,
		Data:     nil,
	}
	c.JSON(http.StatusOK, response)
}
