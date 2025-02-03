package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/models"
)

// LocationRequest struct to capture incoming JSON
type LocationRequest struct {
    Name string `json:"name" binding:"required"`
}

func CreateLocation(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req LocationRequest
        // Bind JSON to struct and check for errors
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
            return
        }

        // Create location model based on request
        newLocation := models.Location{
            Name: req.Name,
        }

        // Save the location to the database
        if err := db.Create(&newLocation).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create location"})
            return
        }

        // Send a response back indicating the location was created
        c.JSON(http.StatusCreated, gin.H{
            "message": "Location created successfully",
            "location": newLocation,
        })
    }
}

func FetchLocation(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var locations []models.Location
        if err := db.Raw("SELECT * FROM locations").Scan(&locations).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, locations)
    }
}
