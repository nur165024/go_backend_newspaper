package setup

import (
	settingsApplication "gin-quickstart/internal/settings/application"
	settingsInfrastructure "gin-quickstart/internal/settings/infrastructure"
	settingsInterfaces "gin-quickstart/internal/settings/interfaces"
	"gin-quickstart/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupSettingsModule(db *sqlx.DB, router *gin.Engine) {
	settingsRepo := settingsInfrastructure.NewSettingsRepository(db)
	settingServices := settingsApplication.NewSettingsServices(settingsRepo)
	settingsHandler := settingsInterfaces.NewSettingsHandler(settingServices)
	
	setupSettingsRoutes(router, settingsHandler)
}

func setupSettingsRoutes(router *gin.Engine, settingsHandler *settingsInterfaces.SettingsHandler) {
	// Load JWT config
	authMiddleware := middleware.GetAuthMiddleware()

	settingsGroup := router.Group("/api/v1/settings")

	// public routes
	settingsGroup.GET("/", settingsHandler.GetAllSettings)
	settingsGroup.GET("/:id", settingsHandler.GetSettingsByID)

	// protected routes
	protected  := settingsGroup.Use(authMiddleware)
	{
		protected.POST("/", settingsHandler.CreateSettings)
		protected.PUT("/:id", settingsHandler.UpdateSettings)
		protected.DELETE("/:id", settingsHandler.DeleteSettings)
	}
}