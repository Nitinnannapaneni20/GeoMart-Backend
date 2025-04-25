package controllers

import (
    "net/http"
    "GeoMart-Backend/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/golang-jwt/jwt/v4"
)

// Renamed helper to avoid conflict
func extractSubFromOrder(c *gin.Context) (string, bool) {
    userToken, exists := c.Get("user")
    if !exists {
        return "", false
    }

    jwtToken, ok := userToken.(*jwt.Token)
    if !ok || jwtToken == nil {
        return "", false
    }

    claims, ok := jwtToken.Claims.(jwt.MapClaims)
    if !ok {
        return "", false
    }

    sub, ok := claims["sub"].(string)
    return sub, ok
}

// POST /api/order/save
func SaveOrder(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        sub, ok := extractSubFromOrder(c)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        var order models.Order

        if err := c.ShouldBindJSON(&order); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        if order.TransactionID == "" || len(order.Items) == 0 || order.TotalAmount == 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
            return
        }

        order.UserID = sub
        if err := db.Create(&order).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order", "details": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Order saved", "order_id": order.ID})
    }
}

// GET /api/orders/user
func GetOrdersByUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        sub, ok := extractSubFromOrder(c)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        var orders []models.Order
        if err := db.Where("user_id = ?", sub).Order("created_at DESC").Find(&orders).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders", "details": err.Error()})
            return
        }

        c.JSON(http.StatusOK, orders)
    }
}
