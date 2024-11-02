package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
        "GeoMart-Backend/middleware"

)

// ProductTypeRoutes defines routes related to product type operations
func ProductTypeRoutes(router *gin.Engine, db *gorm.DB) {
    router.POST("/api/create/product-types",  middleware.JWTMiddleware(), controllers.CreateProductType(db))
}
