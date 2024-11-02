package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/middleware"

)

// LocationRoutes defines routes related to Location operations
func LocationRoutes(router *gin.Engine, db *gorm.DB) {
    router.POST("/api/create/location",  middleware.JWTMiddleware(), controllers.CreateLocation(db))

    // Route for fetching all locations
    router.GET("/api/locations", middleware.JWTMiddleware(), controllers.FetchLocation(db))
}