package setup

import (
	userApplication "gin-quickstart/internal/user/application"
	userInfrastructure "gin-quickstart/internal/user/infrastructure"
	userInterfaces "gin-quickstart/internal/user/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupUserModule(db *sqlx.DB, router *gin.Engine) {
	userRepo := userInfrastructure.NewUserRepository(db)
	userServices := userApplication.NewUserServices(userRepo)
	userHandler := userInterfaces.NewUserHandler(userServices)
	
	setupUserRoutes(router, userHandler)
}

func setupUserRoutes(router *gin.Engine, userHandler *userInterfaces.UserHandler) {
	userGroup := router.Group("/api/v1/users")
	{
		userGroup.GET("/", userHandler.GetAllUsers)
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.POST("/login", userHandler.LoginUser)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}
	
	router.GET("/api/v1/users/email/:email", userHandler.GetUserByEmail)
}
