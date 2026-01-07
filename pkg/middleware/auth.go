package middleware

import (
	"fmt"
	"gin-quickstart/pkg/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtServices *auth.JWTSecret) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		fmt.Printf("🔑 Middleware secret check\n")
		
		claims, err := jwtServices.ValidateToken(tokenString)

		if err != nil {
			fmt.Printf("❌ Token validation error: %v\n", err)

			// Check if token is expired
			if strings.Contains(err.Error(), "token is expired") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired, please login again"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			}
			c.Abort()
			return
		}

		fmt.Printf("✅ Token valid for user: %s\n", claims.Email)

		c.Set("id", claims.ID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Set("user_name", claims.UserName)
		c.Next()
	}
}