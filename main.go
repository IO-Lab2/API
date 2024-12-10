package main

import (
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"time"

	main_router "io-project-api/routes"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logging.Logger.Info("Starting server...")
	if _, err := database.InitDB(); err != nil {
		logging.Logger.Fatal("Failed to connect to database:", zap.Error(err))
	}

	defer logging.Sync()

	router := gin.New()
	router.Use(ginzap.Ginzap(logging.LoggerRaw, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logging.LoggerRaw, true))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		ExposeHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		MaxAge:           12 * time.Hour,
	}))

	huma_config := huma.DefaultConfig("IO Project API", "v0")

	api := humagin.New(router, huma_config)
	main_router.RegisterAPIRoutes(api, "/api")

	logging.Logger.Info("Server started")

	// Start the server
	if err := router.Run("0.0.0.0:8000"); err != nil {
		logging.Logger.Fatal("Failed to start server:", zap.Error(err))
	}
}
