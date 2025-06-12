package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Log struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Timestamp   string `gorm:"not null" json:"timestamp"`
	Method      string `gorm:"size:10;not null" json:"method"`
	Path        string `gorm:"size:255;not null" json:"path"`
	ClientIP    string `gorm:"size:100" json:"client_ip"`
	StatusCode  int    `gorm:"not null" json:"status_code"`
	AuthToken   string `gorm:"type:text" json:"auth_token,omitempty"`
	QueryParams string `gorm:"type:text" json:"query_params,omitempty"`
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
	DB.AutoMigrate(&Log{})
	fmt.Println("✅ Database connection established successfully")
}

// MIDDLEWARE: LOGGING
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		logEntry := Log{
			Timestamp:   time.Now().Format(time.RFC3339),
			Method:      c.Request.Method,
			Path:        c.Request.URL.Path,
			ClientIP:    c.ClientIP(),
			StatusCode:  c.Writer.Status(),
			AuthToken:   c.GetHeader("Authorization"),
			QueryParams: c.Request.URL.RawQuery,
		}

		if err := DB.Create(&logEntry).Error; err != nil {
			log.Println("Failed to save log to database:", err)
		} else {
			log.Println("Log saved to database.")
		}

		// Print to terminal
		logJSON, _ := json.MarshalIndent(logEntry, "", "  ")
		fmt.Println(string(logJSON))
	}
}
//auth
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		token := parts[1]
		if token != "mysecrettoken123" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Next()
	}
}
func main() {
	Init()
	router := gin.Default()
	router.Use(LoggingMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/first", func(c *gin.Context) {
		c.String(200, "First Handler")
	})

	// Protected group
	protected := router.Group("/", AuthMiddleware())
	{
		protected.GET("/second", func(c *gin.Context) {
			c.String(200, "Second Handler")
		})
		protected.GET("/third", func(c *gin.Context) {
			c.String(200, "Third Handler with another function")
		})
	}

	log.Println(" Server starting at https://localhost:3000")

	// Start HTTPS server
	err := router.RunTLS(":3000", "cert.pem", "key.pem")
	if err != nil {
		log.Fatalf(" Error starting HTTPS server: %v", err)
	}
}
