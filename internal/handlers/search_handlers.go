package handlers

import (
	"context"
	"io-project-api/internal/models"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func SearchHandler(ctx context.Context, input *models.SearchInput) (*responses.ScientistsResponse, error) {

	response := &responses.ScientistsResponse{}
	result, err := services.SearchForScientists(input)
	if len(result) == 0 || err != nil {
		return nil, huma.Error400BadRequest("Failed to search for scientists")
	}

	response.Body = result
	return response, nil
}
