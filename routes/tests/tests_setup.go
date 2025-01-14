package tests

import (
	"io-project-api/internal/database"
	internal "io-project-api/routes"
	"log"
	"path/filepath"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func TestSetUP() *gin.Engine {
	//Załadowanie pliku .env
	envPath := filepath.Join("..", "..", ".env")
	log.Println("Loading .env file from: ", envPath)
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
	//Inicjalizacja bazy danych
	if _, err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database initialized")

	//Konfiguracja routera
	router := gin.New()
	config := huma.DefaultConfig("test API", "v1")
	api := humagin.New(router, config)
	internal.RegisterAPIRoutes(api, "api")

	return router
}
