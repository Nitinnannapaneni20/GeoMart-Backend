package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/middleware" // Import the middleware package
)

// UserRoutes defines routes related to user operations
func UserRoutes(router *gin.Engine, db *gorm.DB) {
    // Protect the /api/user_data route with JWT middleware
    router.GET("/api/user_data", middleware.JWTMiddleware(), controllers.GetUserData(db))
}
