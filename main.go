package main

import (
	"github.com/gin-gonic/gin"
)

type Message struct {
	Sentence string `json:"sentence"`
}

func message(c *gin.Context) {
	c.JSON(200, &Message{"This is a test message"})
}

func main() {
	const port = "localhost:8080"
	r := gin.Default()
	r.GET("/", message)
	r.Run(port)
}
