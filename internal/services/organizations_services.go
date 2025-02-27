package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/repositories"
	"io-project-api/internal/requests"
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
func CreateOrganization(input *requests.CreateOrganization) (uuid.UUID, error) {
	id := uuid.New()
	err := repositories.CreateOrganization(database.GetDB(), id, input)
	if err != nil {
		logging.Logger.Error("Error: Failed to create organization!", err)
		return uuid.Nil, err
	}
	return id, nil
}
func UpdateOrganization(input *requests.UpdateOrganization) (*responses.UpdateOrganizationResponse, error) {
	err := repositories.UpdateOrganization(database.GetDB(), input)
	if err != nil {
		logging.Logger.Error("Error: Failed to update organization")
		return nil, err
	}
	updatedOrganization, err := repositories.OrganizationByID(database.GetDB(), input.ID)
	if err != nil {
		logging.Logger.Error("Error: Failed to retrieve updated organization")
		return nil, err
	}
	return &responses.UpdateOrganizationResponse{Body: *updatedOrganization}, nil
}
func DeleteOrganization(input *requests.DeleteOrganization) error {
	err := repositories.DeleteOrganization(database.GetDB(), input)
	if err != nil {
		logging.Logger.Error("Error: Failed to delete organization")
		return err
	}
	return nil
}
