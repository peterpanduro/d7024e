package main

import (
	"fmt"
	"d7024e/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting node")
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.Run()
}
