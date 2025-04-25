package routes

import (
    "GeoMart-Backend/controllers"
    "GeoMart-Backend/middleware"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func OrderRoutes(router *gin.Engine, db *gorm.DB) {
    router.POST("/api/order/save", middleware.JWTMiddleware(), controllers.SaveOrder(db))
    router.GET("/api/orders/user", middleware.JWTMiddleware(), controllers.GetOrdersByUser(db))
}
