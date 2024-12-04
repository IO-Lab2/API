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
	ErrScientistOrganizationNotFound = errors.New("scientist organization not found for the given ID")
)

func GetScientistOrganizationByID(id uuid.UUID) ([]responses.ScientistOrganizationBody, error) {
	logging.Logger.Info("INFO: Retrieving scientist organization by ID")
	scientistOrganization, err := repositories.ScientistOrganizationByID(database.GetDB(), id)
	if err != nil {
		logging.Logger.Info("ERROR: Error querying Scientist Organization by ID", zap.Error(err))
		return nil, err
	}

	if len(scientistOrganization) == 0 {
		logging.Logger.Warn("WARN: No Scientist Organization found", zap.String("ID", id.String()))

		return nil, ErrScientistOrganizationNotFound
	}
	logging.Logger.Info("INFO: Successfully retrieving scientist organization by ID")

	return scientistOrganization, nil
}
