// pkg/middleware/xss_protection.go
package middleware

import (
	"html"
	"strings"

	"github.com/gin-gonic/gin"
)

func XSSProtectionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request body
		var body map[string]interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.Next()
			return
		}

		// Sanitize all string fields
		sanitizeMap(body)

		// Set sanitized body back
		c.Set("sanitized_body", body)
		c.Next()
	}
}

func sanitizeMap(data map[string]interface{}) {
	for key, value := range data {
		switch v := value.(type) {
		case string:
			data[key] = sanitizeString(v)
		case map[string]interface{}:
			sanitizeMap(v)
		}
	}
}

func sanitizeString(input string) string {
	// HTML escape
	escaped := html.EscapeString(input)
	
	// Remove dangerous patterns
	dangerous := []string{
		"<script", "</script>", "javascript:", "onload=", "onerror=", 
		"onclick=", "onmouseover=", "<iframe", "</iframe>",
	}
	
	for _, pattern := range dangerous {
		escaped = strings.ReplaceAll(strings.ToLower(escaped), pattern, "")
	}
	
	return escaped
}
