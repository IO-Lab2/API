package services

import (
	"errors"
	"io-project-api/internal/database"
	"io-project-api/internal/repositories"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrScientistNotFound = errors.New("scientist not found for the given ID")
)

func GetScientistByID(id uuid.UUID) ([]responses.ScientistBody, error) {
	scientist, err := repositories.ScientistByID(database.GetDB().DB, id)
	if err != nil {
		zap.L().Error("Error querying Scientist by ID", zap.Error(err))
		return nil, err
	}

	if len(scientist) == 0 {
		zap.L().Warn("No Scientist found", zap.String("ID", id.String()))
		return nil, ErrScientistNotFound
	}

	return scientist, nil
}
