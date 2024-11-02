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

        // Create the product directly using category_id and type_id
        product := models.ProductData{
            Name:        req.Name,
            Brand:       req.Brand,
            Quantity:    req.Quantity,
            CategoryID:  req.CategoryID,
            TypeID:      req.TypeID,
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
