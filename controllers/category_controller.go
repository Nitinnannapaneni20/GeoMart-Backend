package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/models"
)

type CreateCategoryRequest struct {
    Name string `json:"name" binding:"required"`
}

// CreateCategory creates a new category
func CreateCategory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Name       string `json:"name" binding:"required"`
            LocationID uint   `json:"location_id" binding:"required"`
        }

        // Bind the JSON request to the CreateCategoryRequest struct
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Check if the category already exists
        var category models.Category
        if err := db.Where("name = ? AND location_id = ?", req.Name, req.LocationID).First(&category).Error; err == nil {
            c.JSON(http.StatusConflict, gin.H{"error": "Category already exists"})
            return
        }

        // Create new category
        category = models.Category{Name: req.Name, LocationID: req.LocationID}
        if err := db.Create(&category).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully", "category": category})
    }
}
