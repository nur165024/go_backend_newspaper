package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupAllModules(db *sqlx.DB, router *gin.Engine) {
	SetupUserModule(db, router)
	SetupCategoryModule(db, router)
}
