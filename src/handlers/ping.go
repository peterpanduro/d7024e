package handlers

import (
	"d7024e/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Ping(c *gin.Context) {
	response := models.PONG
    c.JSON(http.StatusOK, response)
}
