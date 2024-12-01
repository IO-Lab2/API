package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/repositories"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrScientistNotFound = errors.New("scientist not found for the given ID")
)

func GetScientistByID(id uuid.UUID) ([]responses.ScientistBody, error) {
	logging.Logger.Info("INFO: Retrieving scientist by ID")
	scientist, err := repositories.ScientistByID(database.GetDB(), id)
	if err != nil {
		logging.Logger.Error("ERROR: Error while fetching Scientist by ID: ", zap.Error(err))
		return nil, err
	}

	if len(scientist) == 0 {
		logging.Logger.Warn("WARN: No Scientist found", zap.String("ID", id.String()))
		return nil, ErrScientistNotFound
	}
	logging.Logger.Info("INFO: Successfully retrieving scientist by ID")
	return scientist, nil
}
