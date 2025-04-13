package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/middleware" // Import the middleware package
)

// UserRoutes defines routes related to user operations
func UserRoutes(router *gin.Engine, db *gorm.DB) {
    api := router.Group("/api")
    {
        api.GET("/user_data", middleware.JWTMiddleware(), controllers.GetUserData(db))
        api.POST("/user/sync", controllers.SyncUser(db)) // no auth needed, sync from frontend
        api.GET("/user/me", middleware.JWTMiddleware(), controllers.GetProfile(db))
        api.PUT("/user/me", middleware.JWTMiddleware(), controllers.UpdateProfile(db))
    }
}
