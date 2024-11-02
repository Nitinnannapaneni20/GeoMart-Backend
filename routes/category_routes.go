package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
        "GeoMart-Backend/middleware"

)

// CategoryRoutes defines routes related to category operations
func CategoryRoutes(router *gin.Engine, db *gorm.DB) {
    router.POST("/api/create/categories",  middleware.JWTMiddleware(), controllers.CreateCategory(db))
}
