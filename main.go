package main

import (
	"data_mapping/database"
	"data_mapping/handlers"
	"data_mapping/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(middleware.LoggingMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Data Mapping API",
		})
	})
	router.POST("/login", handlers.LoginHandler(database.DB))

	auth := router.Group("/")
	auth.Use(handlers.JWTAuthMiddleware())
	{
		auth.POST("/clients", handlers.CreateClient(database.DB))
		auth.GET("/clients", handlers.ListClients(database.DB))
		auth.DELETE("/clients/delete/:id", handlers.DeleteClient(database.DB))

		auth.POST("/clients/:client_id/mappings", handlers.CreateMappings(database.DB))
		auth.GET("/clients/:client_id/mappings", handlers.GetMappings(database.DB))
		auth.DELETE("/clients/:client_id/mappings/:mapping_id", handlers.DeleteMappings(database.DB))

		auth.POST("/clients/:client_id/transform", handlers.TransformHandler(database.DB))
	}
	fmt.Println("Starting server on :https://localhost:8080")
	router.RunTLS(":8080", "cert.pem", "key.pem")
}
