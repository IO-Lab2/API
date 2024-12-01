package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetScientistPublicationByID(ctx context.Context, input *requests.ScientistPublicationID) (*responses.ScientistPublicationResponse, error) {
	logging.Logger.Info("INFO: Handling GetScientistPublicationByID request")
	response := &responses.ScientistPublicationResponse{}

	resultingScientistsPublications, err := services.GetScientistPublicationByID(input.ID)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to get ScientistPublication by ID:", err)
		return nil, huma.Error400BadRequest("Failed to get ScientistPublication by ID")
	}

	logging.Logger.Info("INFO: Successfully retrieved ScientistPublication by ID")
	response.Body = resultingScientistsPublications
	return response, nil
}
