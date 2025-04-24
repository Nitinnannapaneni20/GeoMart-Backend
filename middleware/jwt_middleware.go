package middleware

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "github.com/MicahParks/keyfunc"
)

// JWKS for fetching Auth0's public keys
var jwks *keyfunc.JWKS

// InitializeJWTMiddleware initializes the JWKS with the given issuer URL
func InitializeJWTMiddleware(issuerURL string) {
    jwksURL := fmt.Sprintf("%s/.well-known/jwks.json", issuerURL)
    var err error

    jwks, err = keyfunc.Get(jwksURL, keyfunc.Options{
        RefreshErrorHandler: func(err error) {
            fmt.Printf("There was an error with the jwt.Keyfunc\nError: %s\n", err.Error())
        },
        RefreshInterval: time.Hour,
    })

    if err != nil {
        panic(fmt.Sprintf("Failed to create JWKS from URL: %s\nError: %s", jwksURL, err.Error()))
    }

    fmt.Println("JWKS initialized successfully.")
}

// JWTMiddleware validates the JWT from the appSession cookie (not from header)
func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // First: try Authorization header
        tokenStr := ""
        authHeader := c.GetHeader("Authorization")
        if authHeader != "" && len(authHeader) > 7 && authHeader[:7] == "Bearer " {
            tokenStr = authHeader[7:]
        }

        // If not found in header, try appSession cookie
        if tokenStr == "" {
            cookie, err := c.Cookie("appSession")
            if err == nil {
                tokenStr = cookie
            }
        }

        if tokenStr == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No token provided"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenStr, jwks.Keyfunc)
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("user", token)
        c.Next()
    }
}
