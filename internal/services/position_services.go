package services

import (
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"
)

func GetPositions() ([]models.PositionFilter, error) {

	result, err := repositories.QueryGetPositions(database.GetDB())
	if err != nil {
		logging.Logger.Error("ERROR: Error getting positions:", err)
		return nil, err
	}

	return result, nil
}
