package cmd

import (
	"gin-quickstart/config"
	// category
	categoryApplication "gin-quickstart/internal/category/application"
	categoryInfrastructure "gin-quickstart/internal/category/infrastructure"
	categoryInterfaces "gin-quickstart/internal/category/interfaces"

	// user
	userApplication "gin-quickstart/internal/user/application"
	userInfrastructure "gin-quickstart/internal/user/infrastructure"
	userInterfaces "gin-quickstart/internal/user/interfaces"
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

	// user layers
	userRepo := userInfrastructure.NewUserRepository(db)
	userServices := userApplication.NewUserServices(userRepo)
	userHandler := userInterfaces.NewUserHandler(userServices)
	
	// Setup routes
	setupUserRoutes(router, userHandler)
	

	// category layers
	categoryRepo := categoryInfrastructure.NewCategoryRepository(db)
	categoryServices := categoryApplication.NewCategoryServices(categoryRepo)
	categoryHandler := categoryInterfaces.NewCategoryHandler(categoryServices)

	// Setup routes
	setupCategoryRoutes(router, categoryHandler)

	// SERVER ADDRESS AND ROUTE
	s := &http.Server{
		Addr: serverPort,
		Handler: router,
	}

	s.ListenAndServe()
}

// USERS ROUTES
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
	
	// Separate route for email (conflicts with :id)
	router.GET("/api/v1/users/email/:email", userHandler.GetUserByEmail)
}

// CATEGORY ROUTES
func setupCategoryRoutes(router *gin.Engine, categoryHandler *categoryInterfaces.CategoryHandler) {
	userGroup := router.Group("/api/v1/categories")
	{
		userGroup.GET("/", categoryHandler.GetAllCategories)
		userGroup.POST("/", categoryHandler.CreateCategory)
		userGroup.GET("/:id", categoryHandler.GetCategoryById)
		userGroup.PUT("/:id", categoryHandler.UpdateCategory)
		userGroup.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}