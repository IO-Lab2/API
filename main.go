package main

import (
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"time"

	main_router "io-project-api/routes"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logging.Logger.Info("Starting server...")
	database.InitDB()

	defer logging.Sync()

	router := gin.New()
	router.Use(ginzap.Ginzap(logging.LoggerRaw, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logging.LoggerRaw, true))

	huma_config := huma.DefaultConfig("IO Project API", "v0")

	api := humagin.New(router, huma_config)
	main_router.RegisterAPIRoutes(api, "/api")

	logging.Logger.Info("Server started")

	// Start the server
	if err := router.Run("0.0.0.0:8000"); err != nil {
		logging.Logger.Fatal("Failed to start server:", zap.Error(err))
	}
}
