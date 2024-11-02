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
    LocationID uint   `json:"location_id" binding:"required"`
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

        // Check if the location exists
        var location models.Location
        if err := db.First(&location, req.LocationID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
            return
        }

        // Check if the category exists within the specified location
        var category models.Category
        if err := db.Where("id = ? AND location_id = ?", req.CategoryID, req.LocationID).First(&category).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found in the specified location"})
            return
        }

        // Check if the product type already exists under the given category and location
        var productType models.ProductType
        if err := db.Where("name = ? AND category_id = ? AND location_id = ?", req.Name, req.CategoryID, req.LocationID).First(&productType).Error; err == nil {
            c.JSON(http.StatusConflict, gin.H{"error": "Product type already exists under this category in the specified location"})
            return
        }

        // Create new product type
        productType = models.ProductType{Name: req.Name, CategoryID: req.CategoryID, LocationID: req.LocationID}
        if err := db.Create(&productType).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product type"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Product type created successfully", "product_type": productType})
    }
}
