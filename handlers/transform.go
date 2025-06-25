package handlers

import (
	"net/http"

	"data_mapping/models"
	"data_mapping/utils"

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
