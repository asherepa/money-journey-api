package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})
	app.Run(":8010")
}
