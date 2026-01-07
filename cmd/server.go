package cmd

import (
	"gin-quickstart/cmd/setup"
	"gin-quickstart/config"
	"log"

	"gin-quickstart/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()
	serverCnf, err := config.GetServerConfig()
	if err != nil {
		log.Fatal("Failed to load server config:", err)
	}
	serverPort := ":" + serverCnf.Port

	// Load database config with error handling
	dbCnf, err := config.GetDatabaseConfig()
	if err != nil {
		log.Fatal("Failed to load database config:", err)
	}

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

	// setup all modules
	setup.SetupAllModules(db, router)

	// SERVER ADDRESS AND ROUTE
	s := &http.Server{
		Addr: serverPort,
		Handler: router,
	}

	s.ListenAndServe()
}