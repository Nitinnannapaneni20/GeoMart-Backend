package controllers

import (
    "net/http"
    "GeoMart-Backend/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SaveOrder(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var order models.Order

        if err := c.ShouldBindJSON(&order); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        // Optional: Validate fields
        if order.TransactionID == "" || len(order.Items) == 0 || order.TotalAmount == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
            return
        }

        // Save order
        if err := db.Create(&order).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order", "details": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Order saved", "order_id": order.ID})
    }
}
