package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/models"
)

type CreateProductRequest struct {
    CategoryName  string  `json:"category_name" binding:"required"`
    ProductType   string  `json:"product_type" binding:"required"`
    Name          string  `json:"name" binding:"required"`
    Brand         string  `json:"brand" binding:"required"`
    Quantity      uint    `json:"quantity" binding:"required"`
    Cost          float64 `json:"cost" binding:"required"`
    Description   string  `json:"description" binding:"required"`
    ImageURL      string  `json:"image_url"`
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req CreateProductRequest

        // Bind the JSON request to the CreateProductRequest struct
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Check if the category exists, create if it doesn't
        var category models.Category
        if err := db.Where("name = ?", req.CategoryName).First(&category).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                // Create new category
                category = models.Category{Name: req.CategoryName}
                if err := db.Create(&category).Error; err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
                    return
                }
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying category"})
                return
            }
        }

        // Check if the product type exists, create if it doesn't
        var productType models.ProductType
        if err := db.Where("name = ? AND category_id = ?", req.ProductType, category.ID).First(&productType).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                // Create new product type
                productType = models.ProductType{Name: req.ProductType, CategoryID: category.ID}
                if err := db.Create(&productType).Error; err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product type"})
                    return
                }
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying product type"})
                return
            }
        }

        // Create the product
        product := models.ProductData{
            Name:        req.Name,
            Brand:       req.Brand,
            Quantity:    req.Quantity,
            CategoryID:  category.ID,
            TypeID:      productType.ID,
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
