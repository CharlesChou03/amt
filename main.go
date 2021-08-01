package main

import (
	"github.com/CharlesChou03/_git/amt.git/config"
	_ "github.com/CharlesChou03/_git/amt.git/docs"
	"github.com/CharlesChou03/_git/amt.git/internal/db"
	"github.com/CharlesChou03/_git/amt.git/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setup() {
	config.Setup()
	db.MySQLDB = db.SetupMySQLDB()
	db.RedisDB = db.SetupRedisDB()
}

func setupRouter() *gin.Engine {
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", handlers.HealthHandler)
	r.GET("/version", handlers.VersionHandler)

	r.GET("/api/tutor/:tutor", handlers.GetTutorHandler)

	return r
}

// @title Swagger
// @version 0.0.1
func main() {
	setup()
	defer db.MySQLDB.Close()
	defer db.RedisDB.Close()
	r := setupRouter()
	r.Run(":9999")
}
