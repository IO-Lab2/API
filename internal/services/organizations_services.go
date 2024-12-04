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
	ErrOrganizationNotFound = errors.New("organization not found for the given ID")
)

func GetOrganizationByID(id uuid.UUID) (*responses.OrganizationBodyExtended, error) {
	logging.Logger.Info("INFO: Retrieving organization by ID")
	organization, err := repositories.OrganizationByID(database.GetDB(), id)
	if err != nil {
		logging.Logger.Error("ERROR: Error querying organization by ID ", zap.Error(err))
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieving organization by ID")
	return organization, nil
}

func GetOrganizationsByScientistID(id uuid.UUID) ([]responses.OrganizationBodyExtended, error) {
	logging.Logger.Info("INFO: Retrieving organization by scientist ID")
	organizations, err := repositories.OrganizationsByScientistID(database.GetDB(), id)
	if err != nil {
		logging.Logger.Error("ERROR: Error querying organizations by scientist ID ", zap.Error(err))
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieving organization by scientist ID")
	return organizations, nil
}

func GetOrganizations() (*responses.ListOfOrganizations, error) {
	logging.Logger.Info("INFO: Retrieving organization")
	organizations, err := repositories.Organizations(database.GetDB())
	if err != nil {
		logging.Logger.Error("ERROR: Error querying organizations", zap.Error(err))
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieving organization")
	return organizations, nil
}
