package services

import (
	"database/sql"
	"errors"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrScientistOrganizationNotFound = errors.New("scientist organization not found for the given ID")
)

func GetScientistOragnizationByID(db *sql.DB, id uuid.UUID) ([]models.ScientistOrganization, error) {
	scientistOrganization, err := repositories.ScientistOragnizationByID(db, id)
	if err != nil {
		zap.L().Error("Error querying Scientist Organization by ID", zap.Error(err))
		return nil, err
	}

	if len(scientistOrganization) == 0 {
		zap.L().Warn("No Scientist Organization found", zap.String("ID", id.String()))
		return nil, ErrScientistOrganizationNotFound
	}

	return scientistOrganization, nil
}
