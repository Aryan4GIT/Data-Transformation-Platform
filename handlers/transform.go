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
