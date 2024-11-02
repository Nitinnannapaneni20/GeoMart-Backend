package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "GeoMart-Backend/middleware"
)

// ProductRoutes defines routes related to product operations
func ProductRoutes(router *gin.Engine, db *gorm.DB) {
    router.POST("/api/create/products",  middleware.JWTMiddleware(), controllers.CreateProduct(db))
}
