package setup

import (
	"gin-quickstart/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupAllModules(db *sqlx.DB, router *gin.Engine) {
	middleware.SetupGlobalMiddleWare(router)
	SetupUserModule(db, router)
	SetupCategoryModule(db, router)
}
