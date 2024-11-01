package routes

import (
    "GeoMart-Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// UserRoutes defines routes related to user operations
func UserRoutes(router *gin.Engine, db *gorm.DB) {
    router.GET("/api/user_data", controllers.GetUserData(db))
}
