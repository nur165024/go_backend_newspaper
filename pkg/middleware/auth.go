package middleware

import (
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

		// Proper Bearer token extraction
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := authHeader[7:] // Remove "Bearer " prefix
		
		claims, err := jwtServices.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("id", claims.ID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Set("user_name", claims.UserName)
		c.Next()
	}
}