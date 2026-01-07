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
	
	// 5. Rate Limiting 
	router.Use(RateLimitMiddleware())
	
	// 6. XSS Protection
	router.Use(XSSProtectionMiddleware())
}


func GetAuthMiddleware() gin.HandlerFunc {
	jwtConfig, err := config.GetJWTConfig()

	if err != nil {
		log.Fatal("Failed to load JWT config: ", err)
	}

	jwtSecret := auth.NewJWTServices(jwtConfig.SecretKey)
	return AuthMiddleware(jwtSecret)
}
