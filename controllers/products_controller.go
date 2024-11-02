package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/models"
)

type CreateProductRequest struct {
    CategoryID   uint    `json:"category_id" binding:"required"`
    TypeID       uint    `json:"type_id" binding:"required"`
    LocationID   uint    `json:"location_id" binding:"required"`
    Name         string  `json:"name" binding:"required"`
    Brand        string  `json:"brand" binding:"required"`
    Quantity     uint    `json:"quantity" binding:"required"`
    Cost         float64 `json:"cost" binding:"required"`
    Description  string  `json:"description" binding:"required"`
    ImageURL     string  `json:"image_url"`
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req CreateProductRequest

        // Bind the JSON request to the CreateProductRequest struct
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Validate the location by checking if it exists
        var location models.Location
        if err := db.First(&location, req.LocationID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
            return
        }

        // Create the product directly using category_id, type_id, and location_id
        product := models.ProductData{
            Name:        req.Name,
            Brand:       req.Brand,
            Quantity:    req.Quantity,
            CategoryID:  req.CategoryID,
            TypeID:      req.TypeID,
            LocationID:  req.LocationID,
            Cost:        req.Cost,
            Description: req.Description,
            ImageURL:    req.ImageURL,
        }

        if err := db.Create(&product).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
    }
}
