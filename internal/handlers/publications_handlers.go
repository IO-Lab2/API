package handlers

import (
	"context"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetPublicationByID(ctx context.Context, input *requests.PublicationID) (*responses.PublicationsResponse, error) {
	response := &responses.PublicationsResponse{}
	resultingPublications, err := services.GetPublicationByID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to get publication by ID")
	}
	response.Body = resultingPublications
	return response, nil
}
