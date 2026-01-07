package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// check for validation error
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			if validationErrors, ok := err.Err.(validator.ValidationErrors); ok {
				errorMessages := make(map[string]string)
				
				for _, fieldError := range validationErrors {
					field := strings.ToLower(fieldError.Field())
					
					switch fieldError.Tag() {
					case "required":
						errorMessages[field] = field + " is required"
					case "email":
						errorMessages[field] = "Invalid email format"
					case "min":
						errorMessages[field] = field + " must be at least " + fieldError.Param() + " characters"
					case "max":
						errorMessages[field] = field + " must not exceed " + fieldError.Param() + " characters"
					case "url":
						errorMessages[field] = "Invalid URL format"
					case "alphanum":
						errorMessages[field] = field + " must contain only letters and numbers"
					default:
						errorMessages[field] = field + " is invalid"
					}
				}
				
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Validation failed",
					"details": errorMessages,
				})
				c.Abort()
				return
			}
		}
	}
}