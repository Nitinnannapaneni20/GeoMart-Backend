package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/middleware" // Import the middleware package
)

// UserRoutes defines routes related to user operations
func UserRoutes(router *gin.Engine, db *gorm.DB) {

    router.GET("/api/user_data", middleware.JWTMiddleware(), controllers.GetUserData(db))
    router.POST("/api/profile/create-if-not-exist", middleware.JWTMiddleware(), controllers.CreateUserIfNotExists(db))
    router.GET("/api/profile/get", middleware.JWTMiddleware(), controllers.GetUserBySub(db))
    router.POST("/api/profile/update", middleware.JWTMiddleware(), controllers.UpdateUserProfile(db))
}