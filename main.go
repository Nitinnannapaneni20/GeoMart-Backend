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
    "time"
//     "GeoMart-Backend/models" // Make sure to import your models package
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Get the issuer URL from the environment variables
    issuerURL := os.Getenv("AUTH0_ISSUER_BASE_URL")
    if issuerURL == "" {
        log.Fatal("AUTH0_ISSUER_BASE_URL is not set in the environment variables")
    }

    // Initialize the JWT Middleware with your Auth0 issuer URL
    middleware.InitializeJWTMiddleware(issuerURL)
    log.Println("JWT Middleware initialized successfully.")

    // Database connection setup
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

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
                PrepareStmt: false,
    })
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Automatically create tables based on the models
    //     err = db.AutoMigrate(&models.Location{})
    //     if err != nil {
    //         log.Fatalf("Failed to migrate database: %v", err)
    //     }
    //     log.Println("Database migration completed successfully.")

    // Configure connection pool settings
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Failed to get DB from GORM: %v", err)
    }
    sqlDB.SetMaxIdleConns(10)         // Set the maximum number of idle connections
    sqlDB.SetMaxOpenConns(100)        // Set the maximum number of open connections
    sqlDB.SetConnMaxLifetime(time.Hour) // Set the maximum lifetime of a connection

    log.Println("Database connection established successfully.")

    // Setup routes
    router := gin.Default()
    routes.UserRoutes(router, db)
    routes.CategoryRoutes(router, db)
    routes.ProductTypeRoutes(router, db)
    routes.ProductRoutes(router, db)
    routes.LocationRoutes(router, db)

    // Start the server
    log.Println("Starting server on port :8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
