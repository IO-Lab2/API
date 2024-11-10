package handlers

import (
	"context"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetBibliometricByID(ctx context.Context, input *requests.BibliometricsID) (*responses.BibliometricsResponse, error) {
	response := &responses.BibliometricsResponse{}
	resultingBibliometrics, err := services.GetBibliometricByID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to get bibliometrics by ID")
	}
	response.Body = resultingBibliometrics
	return response, nil
}
