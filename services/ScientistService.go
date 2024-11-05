package services

import (
	"database/sql"
	"errors"
	"io-project-api/models"
	"io-project-api/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrScientistNotFound = errors.New("scientist not found for the given ID")
)

func GetScientistByID(db *sql.DB, id uuid.UUID) ([]models.Scientist, error) {
	scientist, err := repositories.ScientistByID(db, id)
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
