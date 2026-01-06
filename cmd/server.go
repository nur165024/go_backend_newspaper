package cmd

import (
	"gin-quickstart/config"
	"gin-quickstart/internal/user/application"
	"gin-quickstart/internal/user/infrastructure"
	"gin-quickstart/internal/user/interfaces"
	"gin-quickstart/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()
	serverCnf := config.GetServerConfig()
	serverPort := ":" + serverCnf.Port
	dbCnf := config.GetDatabaseConfig()

	// database connection
	dbConnection := &database.DatabaseConfig {
		Host: dbCnf.DBHost,
		Port: dbCnf.DBPort,
		User: dbCnf.DBUser,
		Password: dbCnf.DBPass,
		DbName: dbCnf.DBName,
	}
	db := database.NewDatabaseConnection(dbConnection)

	// initialization route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": serverCnf.ServerName,
			"status":  http.StatusOK,
			"version": serverCnf.Version,
			"author":  "Nure Alam",
		})
	})


	// Initialize layers
	userRepo := infrastructure.NewPostgresUserRepository(db)
	userService := application.NewUserService(userRepo)
	userHandler := interfaces.NewUserHandler(userService)

	// Setup routes
	setupUserRoutes(router, userHandler)

	// SERVER ADDRESS AND ROUTE
	s := &http.Server{
		Addr: serverPort,
		Handler: router,
	}

	s.ListenAndServe()
}

// USERS ROUTES
func setupUserRoutes(router *gin.Engine, userHandler *interfaces.UserHandler) {
	userGroup := router.Group("/api/v1/users")
	{
		userGroup.GET("/", userHandler.GetAllUsers)
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.POST("/login", userHandler.LoginUser)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}
	// Separate route for email (conflicts with :id)
	router.GET("/api/v1/users/email/:email", userHandler.GetUserByEmail)
}