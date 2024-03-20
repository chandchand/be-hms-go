package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// const secretKey = "aptx4869"

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        fmt.Println("tokenstring: ", tokenString)

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("SECRET_KEY")), nil
        })

        fmt.Println("tokenparse: ", token)

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            return
        }

        userID := claims["userID"].(string)
        firstName := claims["first_name"].(string)
        username := claims["username"].(string)
        role := claims["role"].(string)

        c.Set("userID", userID)
        c.Set("firstName", firstName)
        c.Set("username", username)
        c.Set("role", role)

        c.Next()
    }
}
