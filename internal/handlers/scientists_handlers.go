package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetScientistByID(ctx context.Context, input *requests.ScientistID) (*responses.ScientistResponse, error) {
	logging.Logger.Info("INFO: Handling GetScientistByID request")
	response := &responses.ScientistResponse{}

	resultingScientists, err := services.GetScientistByID(input.ID)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to get scientist by ID:", err)
		return nil, huma.Error400BadRequest("Failed to get scientist by ID")
	}

	logging.Logger.Info("INFO: Successfully retrieved scientist by ID")
	response.Body = resultingScientists
	return response, nil
}
