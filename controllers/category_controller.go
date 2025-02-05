package controllers

import (
    "context"
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

// GetProductsDataByLocation fetches all categories, product types, and products based on location_id
func GetProductsDataByLocation(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        locationID := c.Query("location_id")

        if locationID == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "location_id query parameter is required"})
            return
        }

        var categories []models.Category
        var productTypes []models.ProductType
        var products []models.ProductData

        // **KEY FIX: Use a fresh DB session**
        ctx := context.Background()
        tx := db.WithContext(ctx).Session(&gorm.Session{SkipDefaultTransaction: true}) // Fresh DB session

        // Fetch categories by location_id
        if err := tx.Where("location_id = ?", locationID).Find(&categories).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories", "details": err.Error()})
            return
        }

        // Fetch product types by location_id
        if err := tx.Where("location_id = ?", locationID).Find(&productTypes).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product types", "details": err.Error()})
            return
        }

        // Fetch products by location_id
        if err := tx.Where("location_id = ?", locationID).Preload("Category").Preload("ProductType").Find(&products).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products", "details": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "categories":   categories,
            "productTypes": productTypes,
            "products":     products,
        })
    }
}
