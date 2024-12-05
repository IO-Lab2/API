package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetOrganizationById(ctx context.Context, input *requests.OrganizationID) (*responses.OrganizationResponse, error) {
	logging.Logger.Info("INFO: Handling GetOrganizationById request")
	response := &responses.OrganizationResponse{}
	resultingOrganizations, err := services.GetOrganizationByID(input.ID)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to get organization by ID:", err)
		return nil, huma.Error400BadRequest("Failed to get organization by ID")
	}
	logging.Logger.Info("INFO: Successfully retrieved organization by ID")
	response.Body = resultingOrganizations
	return response, nil
}

func GetOrganizationsByScientistId(ctx context.Context, input *requests.ScientistID) (*responses.ListOfOrganizationsResponse, error) {
	logging.Logger.Info("INFO: Handling GetOrganizationsByScientistId request")
	response := &responses.ListOfOrganizationsResponse{}
	resultingOrganizations, err := services.GetOrganizationsByScientistID(input.ID)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to get organizations by scientist ID:", err)
		return nil, huma.Error400BadRequest("Failed to get organizations by scientist ID")
	}
	logging.Logger.Info("INFO: Successfully retrieved organizations by scientist ID")
	response.Body = resultingOrganizations
	return response, nil
}

func GetOrganizationsHandler(ctx context.Context) (*responses.ListOfOrganizations, error) {
	logging.Logger.Info("INFO: Handling GetOrganizationsHandler request")
	response := &responses.ListOfOrganizations{}
	resultingOrganizations, err := services.GetOrganizations()
	if err != nil {
		logging.Logger.Error("ERROR: Failed to get organizations:", err)
		return nil, huma.Error400BadRequest("Failed to get organizations")
	}
	logging.Logger.Info("INFO: Successfully retrieved organizations")
	response.Body = resultingOrganizations.Body
	return response, nil
}
func CreateOrganization(ctx context.Context, input *requests.CreateOrganization) (*responses.CreateOrganizationResponse, error) {
	response := &responses.CreateOrganizationResponse{}
	createdOrganization, err := services.CreateOrganization(input)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to create bibliometric")
	}
	response.Body = responses.CreateOrganization{ID: createdOrganization}
	return response, nil
}
func UpdateOrganization(ctx context.Context, input *requests.UpdateOrganization) (*responses.UpdateOrganizationResponse, error) {
	response := &responses.UpdateOrganizationResponse{}
	updatedOrganization, err := services.UpdateOrganization(input)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to update organization")
	}
	response = updatedOrganization
	return response, nil
}
