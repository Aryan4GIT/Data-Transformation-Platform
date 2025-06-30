package handlers

import (
	"data_mapping/models"
	"data_mapping/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransformHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.Param("client_id")
		var req models.TransformationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var rules []models.MappingRule
		result := db.Where("client_id = ?", clientID).Find(&rules)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		output, err := utils.Transform(req.InputData, rules)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

// StreamTransformHandler streams and transforms large client JSONs in real-time.
func StreamTransformHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.Param("client_id")
		var rules []models.MappingRule
		result := db.Where("client_id = ?", clientID).Find(&rules)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		transformFunc := func(key string, value interface{}) (string, interface{}) {
			// Find rule for this key
			for _, rule := range rules {
				if len(rule.SourcePath) == 1 && rule.SourcePath[0] == key {
					transformedVal, _ := utils.ApplyTransform(value, rule.TransformType)
					return rule.DestinationPath[0], transformedVal
				}
			}
			return key, value // default: no transformation
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		if err := utils.StreamTransformJSON(c.Request.Body, c.Writer, transformFunc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

// UnifiedTransformHandler handles both standard and large payloads for transformation.
func UnifiedTransformHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.Param("client_id")
		var rules []models.MappingRule
		result := db.Where("client_id = ?", clientID).Find(&rules)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load mapping rules", "details": result.Error.Error()})
			return
		}

		// Limit payload size for security (e.g., 10MB)
		if c.Request.ContentLength > 10*1024*1024 {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "Payload too large. Max 10MB allowed."})
			return
		}

		stream := c.GetHeader("X-Stream-Transform") == "true"
		if !stream && c.Request.ContentLength > 0 && c.Request.ContentLength < 5*1024*1024 {
			var input map[string]interface{}
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input", "details": err.Error()})
				return
			}
			output, err := utils.Transform(input, rules)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Transformation failed", "details": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"success": true, "data": output})
			return
		}

		// Streaming mode for large payloads or if header is set
		c.Writer.Header().Set("Content-Type", "application/json")
		if err := utils.StreamTransformJSONWithRules(c.Request.Body, c.Writer, rules); err != nil {
			c.Error(err) // Log error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Streaming transformation failed", "details": err.Error()})
		}
	}
}
