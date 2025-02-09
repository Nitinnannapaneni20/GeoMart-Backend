package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// SpecialsRoutes defines routes related to Special Products
func SpecialsRoutes(router *gin.Engine, db *gorm.DB) {

    router.GET("/api/special-products-data", controllers.GetSpecialsByLocation(db))

}
