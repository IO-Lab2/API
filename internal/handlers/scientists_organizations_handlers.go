package handlers

import (
	"context"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetScientistOrganizationByID(ctx context.Context, input *requests.ScientistOrganizationID) (*responses.ScientistOrganizationResponse, error) {
	response := &responses.ScientistOrganizationResponse{}
	resultingScientistsOrganizations, err := services.GetScientistOrganizationByID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to get scientist organization by ID")
	}
	response.Body = resultingScientistsOrganizations
	return response, nil
}
