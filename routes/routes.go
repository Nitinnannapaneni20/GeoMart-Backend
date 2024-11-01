package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// SetupRoutes sets up all routes for the application
func SetupRoutes(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    UserRoutes(router, db)

    return router
}
