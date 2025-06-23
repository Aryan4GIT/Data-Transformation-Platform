package main

import (
	"fmt"
	"log"

	"third_party_integrations/handler"
	"third_party_integrations/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server started successfully")
	router := gin.Default()
	router.Use(middleware.LoggingMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/first", func(c *gin.Context) {
		c.String(200, "First Handler")
	})

	protected := router.Group("/", middleware.AuthMiddleware())
	{
		protected.GET("/second", func(c *gin.Context) {
			c.String(200, "Second Handler")
		})
		protected.GET("/third", func(c *gin.Context) {
			c.String(200, "Third Handler with another function")
		})
	}

	router.POST("/login", handler.LoginHandler)
	router.POST("/create-user", handler.CreateUserHandler)
	router.DELETE("/delete-user/:id", handler.DeleteUserHandler)


	log.Println("Server starting at https://localhost:3000")
	err := router.RunTLS(":3000", "cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Error starting HTTPS server: %v", err)
	}
}
