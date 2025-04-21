package controllers

import (
    "net/http"
    "GeoMart-Backend/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func GetUserData(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var users []models.UserData
        if err := db.Find(&users).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, users)
    }
}

func CreateUserIfNotExists(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            GivenName   string `json:"given_name"`
            FamilyName  string `json:"family_name"`
            Email       string `json:"email"`
            Picture     string `json:"picture"` // not stored, just available
            Sub         string `json:"sub"`
        }

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }

        var existing models.UserData
        if err := db.Where("auth0_id = ?", req.Sub).First(&existing).Error; err == nil {
            // User exists — return silently
            c.JSON(http.StatusOK, gin.H{"message": "User already exists"})
            return
        }

        user := models.UserData{
            Auth0ID: req.Sub,
            Name:    req.GivenName + " " + req.FamilyName,
            Email:   req.Email,
        }

        if err := db.Create(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "User created"})
    }
}

func GetUserBySub(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Sub string `json:"sub"`
        }

        if err := c.ShouldBindJSON(&req); err != nil || req.Sub == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Missing sub in request body"})
            return
        }

        var user models.UserData
        if err := db.Where("auth0_id = ?", req.Sub).First(&user).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "name":          user.Name,
            "email":         user.Email,
            "phone":         user.Phone,
            "addressLine1":  user.AddressLine1,
            "addressLine2":  user.AddressLine2,
            "city":          user.City,
            "state":         user.State,
            "zip":           user.Zip,
            "sub":           user.Auth0ID,
        })
    }
}

func UpdateUserProfile(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Name         string `json:"name"`
            Email        string `json:"email"`
            Phone        string `json:"phone"`
            AddressLine1 string `json:"addressLine1"`
            AddressLine2 string `json:"addressLine2"`
            City         string `json:"city"`
            State        string `json:"state"`
            Zip          string `json:"zip"`
            Sub          string `json:"sub"`
        }

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }

        var user models.UserData
        if err := db.Where("auth0_id = ?", req.Sub).First(&user).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        // Update individual fields directly
        user.Name = req.Name
        user.Email = req.Email
        user.Phone = req.Phone
        user.AddressLine1 = req.AddressLine1
        user.AddressLine2 = req.AddressLine2
        user.City = req.City
        user.State = req.State
        user.Zip = req.Zip

        if err := db.Save(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "details": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
    }
}
