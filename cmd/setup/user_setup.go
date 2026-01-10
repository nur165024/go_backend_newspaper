package setup

import (
	"gin-quickstart/config"
	authInfrastructure "gin-quickstart/internal/auth/infrastructure"
	userApplication "gin-quickstart/internal/user/application"
	userInfrastructure "gin-quickstart/internal/user/infrastructure"
	userInterfaces "gin-quickstart/internal/user/interfaces"
	"gin-quickstart/pkg/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupUserModule(db *sqlx.DB, router *gin.Engine) {
	jwtCnf, err := config.GetJWTConfig()
	if err != nil {
		log.Fatal("Failed to load JWT config:", err)
	}

	userRepo := userInfrastructure.NewUserRepository(db)
	authRepo := authInfrastructure.NewRefreshTokenRepository(db)
	userServices := userApplication.NewUserServices(userRepo, authRepo, jwtCnf)
	userHandler := userInterfaces.NewUserHandler(userServices)
	
	setupUserRoutes(router, userHandler)
}

func setupUserRoutes(router *gin.Engine, userHandler *userInterfaces.UserHandler) {
	// Load JWT config
	authMiddleware := middleware.GetAuthMiddleware()

	userGroup := router.Group("/api/v1/users")
	// public routes
	userGroup.POST("/", userHandler.CreateUser)
	userGroup.POST("/login", userHandler.LoginUser)
	
	// protected routes
	protected  := userGroup.Use(authMiddleware)
	{
		protected.GET("/", userHandler.GetAllUsers)
		
		protected.GET("/:id", userHandler.GetUserByID)
		protected.PUT("/:id", userHandler.UpdateUser)
		protected.DELETE("/:id", userHandler.DeleteUser)
	}
	
	protected.GET("/api/v1/users/email/:email", userHandler.GetUserByEmail)
}
