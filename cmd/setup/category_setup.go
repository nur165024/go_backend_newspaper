package setup

import (
	categoryApplication "gin-quickstart/internal/category/application"
	categoryInfrastructure "gin-quickstart/internal/category/infrastructure"
	categoryInterfaces "gin-quickstart/internal/category/interfaces"

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
	categoryGroup := router.Group("/api/v1/categories")
	{
		categoryGroup.GET("/", categoryHandler.GetAllCategories)
		categoryGroup.POST("/", categoryHandler.CreateCategory)
		categoryGroup.GET("/:id", categoryHandler.GetCategoryById)
		categoryGroup.PUT("/:id", categoryHandler.UpdateCategory)
		categoryGroup.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}
