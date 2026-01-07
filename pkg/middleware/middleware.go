package middleware

import (
	"gin-quickstart/config"
	"gin-quickstart/pkg/auth"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupGlobalMiddleWare(router *gin.Engine) {
	// rate limit 
	router.Use(RateLimitMiddleware())
	// logger
	router.Use(CustomLogger())
	// recovery
	router.Use(gin.Recovery())
	// cors
	router.Use(CORS())
}

func GetAuthMiddleware() gin.HandlerFunc {
	jwtConfig, err := config.GetJWTConfig()

	if err != nil {
		log.Fatal("Failed to load JWT config: ", err)
	}

	jwtSecret := auth.NewJWTServices(jwtConfig.SecretKey)
	return AuthMiddleware(jwtSecret)
}
