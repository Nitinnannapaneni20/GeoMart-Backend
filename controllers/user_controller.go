package controllers

import (
    "net/http"
    "GeoMart-Backend/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "fmt"
)

func SyncUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var body struct {
            Auth0ID string `json:"auth0Id"`
            Name    string `json:"name"`
            Email   string `json:"email"`
        }

        if err := c.BindJSON(&body); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        if body.Auth0ID == "" || body.Email == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required Auth0 user info"})
            return
        }

        var user models.UserData
        if err := db.Where("auth0_id = ?", body.Auth0ID).First(&user).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                // Create new user
                user = models.UserData{
                    Auth0ID: body.Auth0ID,
                    Name:    body.Name,
                    Email:   body.Email,
                }
                if err := db.Create(&user).Error; err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                    return
                }
                fmt.Println("New user created:", user)
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
        } else {
            fmt.Println("User already exists:", user)
        }

        c.JSON(http.StatusOK, user)
    }
}


func GetProfile(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        auth0ID := c.GetString("auth0Id") // from middleware
        var user models.UserData
        if err := db.Where("auth0_id = ?", auth0ID).First(&user).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusOK, user)
    }
}

func UpdateProfile(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        auth0ID := c.GetString("auth0Id")
        var updates models.UserData

        if err := c.BindJSON(&updates); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }

        var user models.UserData
        if err := db.Where("auth0_id = ?", auth0ID).First(&user).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        // Update profile fields
        user.Name = updates.Name
        user.Phone = updates.Phone
        user.AddressLine1 = updates.AddressLine1
        user.AddressLine2 = updates.AddressLine2
        user.City = updates.City
        user.State = updates.State
        user.Zip = updates.Zip

        if err := db.Save(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, user)
    }
}

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
