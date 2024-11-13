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
	ErrScientistOrganizationNotFound = errors.New("scientist organization not found for the given ID")
)

func GetScientistOrganizationByID(id uuid.UUID) ([]responses.ScientistOrganizationBody, error) {
	scientistOrganization, err := repositories.ScientistOragnizationByID(database.GetDB().DB, id)
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
