package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// CategoryRoutes defines routes related to category operations
func CategoryRoutes(router *gin.Engine, db *gorm.DB) {

    router.GET("/api/special-products-data", controllers.GetSpecialsByLocation(db))

}
