package services

import (
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"
)

func GetPublishers() ([]models.PublisherFilter, error) {

	result, err := repositories.QueryGetPublishers(database.GetDB())
	if err != nil {
		logging.Logger.Error("ERROR: Error getting publishers:", err)
		return nil, err
	}

	return result, nil
}
