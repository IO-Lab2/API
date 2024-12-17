package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func SearchHandler(ctx context.Context, input *models.SearchInput) (*responses.ScientistsResponse, error) {
	logging.Logger.Info("INFO: Handling SearchHandler request")
	response := &responses.ScientistsResponse{}

	result, err := services.SearchForScientists(input)
	if result == nil || len(result) == 0 || err != nil {
		logging.Logger.Error("ERROR: Failed to search for scientists:", err)
		return nil, huma.Error400BadRequest(err.Error())
	}

	logging.Logger.Info("INFO: Successfully retrieved search results for scientists")
	response.Body = result
	return response, nil
}
