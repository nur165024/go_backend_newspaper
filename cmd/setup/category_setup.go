package setup

import (
	categoryApplication "gin-quickstart/internal/category/application"
	categoryInfrastructure "gin-quickstart/internal/category/infrastructure"
	categoryInterfaces "gin-quickstart/internal/category/interfaces"
	"gin-quickstart/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupCategoryModule(db *sqlx.DB, router *gin.Engine) {
	categoryRepo := categoryInfrastructure.NewCategoryRepository(db)
	categoryServices := categoryApplication.NewCategoryServices(categoryRepo)
	categoryHandler := categoryInterfaces.NewCategoryHandler(categoryServices)
	
	setupCategoryRoutes(router, categoryHandler)
}

func setupCategoryRoutes(router *gin.Engine, categoryHandler *categoryInterfaces.CategoryHandler) {
	// Load JWT config
	authMiddleware := middleware.GetAuthMiddleware()
	
	categoryGroup := router.Group("/api/v1/categories")
	// public routes
	categoryGroup.GET("/", categoryHandler.GetAllCategories)
	categoryGroup.GET("/:id", categoryHandler.GetCategoryByID)

	protected  := categoryGroup.Use(authMiddleware)
	{
		protected.POST("/", categoryHandler.CreateCategory)
		protected.PUT("/:id", categoryHandler.UpdateCategory)
		protected.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}
