package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetScientistOrganizationByID(ctx context.Context, input *requests.ScientistOrganizationID) (*responses.ScientistOrganizationResponse, error) {
	logging.Logger.Info("INFO: Handling GetScientistOrganizationByID request")
	response := &responses.ScientistOrganizationResponse{}

	resultingScientistsOrganizations, err := services.GetScientistOrganizationByID(input.ID)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to get scientist organization by ID:", err)
		return nil, huma.Error400BadRequest("Failed to get scientist organization by ID")
	}

	logging.Logger.Info("INFO: Successfully retrieved scientist organization by ID")
	response.Body = resultingScientistsOrganizations
	return response, nil
}
