package handlers

import (
	"net/http"
	"strconv"

	"data_mapping/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMappings(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID, err := strconv.Atoi(c.Param("client_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
			return
		}
		var rules []models.MappingRule
		if err := c.ShouldBindJSON(&rules); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i := range rules {
			rules[i].ClientID = uint(clientID)
		}
		if result := db.Create(&rules); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusCreated, rules)
	}
}

func GetMappings(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.Param("client_id")
		var rules []models.MappingRule
		result := db.Where("client_id = ?", clientID).Find(&rules)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, rules)
	}
}

func DeleteMappings(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.Param("client_id")
		result := db.Where("client_id = ?", clientID).Delete(&models.MappingRule{})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
