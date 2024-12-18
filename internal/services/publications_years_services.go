package services

import (
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"
)

func GetPublicationsYears() ([]models.PublicationsYear, error) {

	years, err := repositories.QueryGetPublicationsYears(database.GetDB())
	if err != nil {
		logging.Logger.Error("ERROR: Failed to retrieve publications years:", err)
		return nil, err
	}

	return years, nil
}
