// pkg/middleware/csrf.go
package middleware

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// CSRF for safe methods
		if c.Request.Method == "GET" || c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" || c.Request.Method == "PATCH" || c.Request.Method == "HEAD" || c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		token := c.GetHeader("X-CSRF-Token")
		sessionToken := c.GetHeader("Authorization")

		if !validateCSRFToken(token, sessionToken) {
			c.JSON(403, gin.H{"error": "Invalid CSRF token"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func validateCSRFToken(token, sessionToken string) bool {
	if token == "" || sessionToken == "" {
		return false
	}

	// Generate expected token based on session
	expected := generateTokenForSession(sessionToken)
	return hmac.Equal([]byte(token), []byte(expected))
}

func generateTokenForSession(sessionToken string) string {
	h := hmac.New(sha256.New, []byte("csrf-secret-key-change-this"))
	h.Write([]byte(sessionToken))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func GenerateCSRFToken(sessionToken string) string {
	return generateTokenForSession(sessionToken)
}

// For backward compatibility
func GenerateRandomCSRFToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
