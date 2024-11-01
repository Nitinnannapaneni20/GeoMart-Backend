package main

import (
    "log"
    "GeoMart-Backend/routes"
    "GeoMart-Backend/middleware"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "fmt"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    issuerURL := os.Getenv("AUTH0_ISSUER_BASE_URL")
    if issuerURL == "" {
        log.Fatal("AUTH0_ISSUER_BASE_URL is not set in environment variables")
    }
    middleware.InitializeJWTMiddleware(issuerURL)

    // Database connection without migration or table check
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSLMODE"),
        os.Getenv("DB_TIMEZONE"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Set up routes and start the server
    router := gin.Default()
    routes.UserRoutes(router, db)
    router.Run(":8080")
}
