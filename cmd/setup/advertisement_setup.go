package setup

import (
	adsApplication "gin-quickstart/internal/advertisement/application"
	adsInfrastructure "gin-quickstart/internal/advertisement/infrastructure"
	adsInterfaces "gin-quickstart/internal/advertisement/interfaces"
	"gin-quickstart/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupAdsModule(db *sqlx.DB, router *gin.Engine) {
	adsRepo := adsInfrastructure.NewAdsRepository(db)
	adsServices := adsApplication.NewAdsServices(adsRepo)
	adsHandler := adsInterfaces.NewAdsHandler(adsServices)

	setupAdsRoutes(router, adsHandler)
}

func setupAdsRoutes(router *gin.Engine, adsHandler *adsInterfaces.AdsHandler) {
	// Load JWT config
	authMiddleware := middleware.GetAuthMiddleware()
	
	adsGroup := router.Group("/api/v1/advertisement")
	// public routes
	adsGroup.GET("/", adsHandler.GetAllAds)
	adsGroup.GET("/:id", adsHandler.GetAdsByID)
	
	protected  := adsGroup.Use(authMiddleware)
	{
		protected.POST("/", adsHandler.CreateAds)
		protected.PUT("/:id", adsHandler.UpdateAds)
		protected.DELETE("/:id", adsHandler.DeleteAds)
	}
}