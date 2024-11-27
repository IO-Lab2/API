package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetScientistByID(ctx context.Context, input *requests.ScientistID) (*responses.ScientistsResponse, error) {

	response := &responses.ScientistsResponse{}
	resultingScientists, err := services.GetScientistByID(input.ID)
	if err != nil {
		logging.Logger.Error("Failed to get scientist by ID: ", err)
		return nil, huma.Error400BadRequest("Failed to get scientist by ID")
	}

	response.Body = resultingScientists
	return response, nil
}
