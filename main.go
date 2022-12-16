package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Sentence string `json:"sentence"`
}

func message(c *gin.Context) {
	c.JSON(200, &Message{"This is a test message"})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
