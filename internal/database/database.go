package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

var db *sqlx.DB

type DBConfig struct {
	Host     string
	User     string
	Port     string
	Database string
	Password string
}

func InitDB() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning! No .env file found!")
	}
	config := DBConfig{
		Host:     os.Getenv("PGHOST"),
		User:     os.Getenv("PGUSER"),
		Port:     os.Getenv("PGPORT"),
		Database: os.Getenv("PGDATABASE"),
		Password: os.Getenv("PGPASSWORD"),
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
		config.Host, config.Port, config.User, config.Database, config.Password)

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed db connection! %v", err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(20 * time.Minute)
	log.Println("DB connection initialized!")

	return db, nil
}
func GetDB() *sqlx.DB {
	if db == nil {
		log.Fatal("Database not initialized. Running InitDB() is required first.")
	}
	return db
}
func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database %v", err)
		}
	}
}
