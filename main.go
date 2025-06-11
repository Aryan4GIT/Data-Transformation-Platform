package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/first", func(c *gin.Context) {
		c.String(200, "First Handler")
	})

	router.GET("/second", func(c *gin.Context) {
		c.String(200, "Second Handler")
	})
	router.GET("/third", func(c *gin.Context) {
		c.String(200, "Third Handler with another function")
	})
	router.RunTLS(":3000", "cert.pem", "key.pem")
	fmt.Println("Server starting at http://localhost:3000")
	err := router.Run("0.0.0.0:3000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
