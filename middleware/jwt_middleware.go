package middleware

import (
    "fmt"
    "net/http"
    "time"
    "os"

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

    // Initialize the JWKS
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

// JWTMiddleware validates the JWT from Authorization header OR secure cookie
func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := ""

        // 1. Check Authorization header
        authHeader := c.GetHeader("Authorization")
        if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
            tokenString = authHeader[7:]
        }

        // 2. If not found, try cookie (used by Auth0 session flow)
        if tokenString == "" {
            cookie, err := c.Cookie("appSession") // 🔁 Change this if your Auth0 cookie is named differently
            if err == nil {
                tokenString = cookie
            }
        }

        // If still no token, reject
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token in header or cookie"})
            c.Abort()
            return
        }

        // Parse and validate the JWT
        token, err := jwt.Parse(tokenString, jwks.Keyfunc)
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Optional audience claim check
        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            audience := os.Getenv("AUTH0_AUDIENCE")
            if audience != "" {
                if audList, ok := claims["aud"].([]interface{}); ok {
                    validAud := false
                    for _, aud := range audList {
                        if aud == audience {
                            validAud = true
                            break
                        }
                    }
                    if !validAud {
                        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid audience"})
                        c.Abort()
                        return
                    }
                }
            }
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }

        c.Set("user", token)
        c.Next()
    }
}
