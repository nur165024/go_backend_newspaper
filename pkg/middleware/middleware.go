package middleware

import (
	"gin-quickstart/config"
	"gin-quickstart/pkg/auth"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupGlobalMiddleWare(router *gin.Engine) {
	// 1. Recovery
	router.Use(gin.Recovery())
	
	// 2. Security Headers 
	router.Use(SecurityHeaderMiddleware())
	
	// 3. CORS
	router.Use(CORS())
	
	// 4. Logger
	router.Use(CustomLogger())
	
	// Disable problematic middlewares for now
	// router.Use(RateLimitMiddleware())
	// router.Use(XSSProtectionMiddleware())
	// router.Use(CSRFMiddleware())  // This blocks POST requests
	// router.Use(ValidationMiddleware())
}


func GetAuthMiddleware() gin.HandlerFunc {
	jwtConfig, err := config.GetJWTConfig()

	if err != nil {
		log.Fatal("Failed to load JWT config: ", err)
	}

	jwtSecret := auth.NewJWTServices(jwtConfig.SecretKey, jwtConfig.AccessTokenExpireMinutes, jwtConfig.RefreshTokenExpireDays)
	return AuthMiddleware(jwtSecret)
}