package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/models"
)

type CreateProductTypeRequest struct {
    Name       string `json:"name" binding:"required"`
    CategoryID uint   `json:"category_id" binding:"required"`
}

// CreateProductType creates a new product type
func CreateProductType(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req CreateProductTypeRequest

        // Bind the JSON request to the CreateProductTypeRequest struct
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Check if the category exists
        var category models.Category
        if err := db.First(&category, req.CategoryID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
            return
        }

        // Check if the product type already exists under the given category
        var productType models.ProductType
        if err := db.Where("name = ? AND category_id = ?", req.Name, req.CategoryID).First(&productType).Error; err == nil {
            c.JSON(http.StatusConflict, gin.H{"error": "Product type already exists under this category"})
            return
        }

        // Create new product type
        productType = models.ProductType{Name: req.Name, CategoryID: req.CategoryID}
        if err := db.Create(&productType).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product type"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Product type created successfully", "product_type": productType})
    }
}
