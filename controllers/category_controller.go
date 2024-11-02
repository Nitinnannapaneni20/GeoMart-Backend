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
        var req CreateCategoryRequest

        // Bind the JSON request to the CreateCategoryRequest struct
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Check if the category already exists
        var category models.Category
        if err := db.Where("name = ?", req.Name).First(&category).Error; err == nil {
            c.JSON(http.StatusConflict, gin.H{"error": "Category already exists"})
            return
        }

        // Create new category
        category = models.Category{Name: req.Name}
        if err := db.Create(&category).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully", "category": category})
    }
}
