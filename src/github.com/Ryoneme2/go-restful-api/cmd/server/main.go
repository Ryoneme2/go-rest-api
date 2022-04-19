package main

import (
	"../../pkg/greeting"
	"github.com/gin-gonic/gin"
)

func main() {
	greet := greeting.Greeting()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": greet,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
