package controllers

import (
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/models"
)

// GetSpecialsByLocation fetches all special products based on location_id
func GetSpecialsByLocation(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        locationID := c.Query("location_id")

        if locationID == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "location_id query parameter is required"})
            return
        }

        var specials []models.Specials

        // **Using a fresh DB session**
        ctx := context.Background()
        tx := db.WithContext(ctx).Session(&gorm.Session{SkipDefaultTransaction: true})

        // Fetch special products by location_id and preload related product details
        if err := tx.Where("location_id = ?", locationID).Preload("Product").Preload("Location").Find(&specials).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch special products", "details": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "specials": specials,
        })
    }
}
