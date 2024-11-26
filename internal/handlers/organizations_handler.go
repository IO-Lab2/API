package handlers

import (
	"context"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetOrganizationById(ctx context.Context, input *requests.OrganizationID) (*responses.OrganizationsResponse, error) {
	response := &responses.OrganizationsResponse{}
	resultingOrganizations, err := services.GetOrganizationByID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to get organization by ID")
	}
	response.Body = resultingOrganizations
	return response, nil
}

func GetOrganizationsByScientistId(ctx context.Context, input *requests.ScientistID) (*responses.ListOfOrganizationsResponse, error) {
	response := &responses.ListOfOrganizationsResponse{}
	resultingOrganizations, err := services.GetOrganizationsByScientistID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to get organizations by scientist ID")
	}
	response.Body = resultingOrganizations
	return response, nil
}
