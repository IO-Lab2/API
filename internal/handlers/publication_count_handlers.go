package handlers

import (
	"context"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetPublicationCountHandler(ctx context.Context) (*responses.PublicationCountResponse, error) {
	response := &responses.PublicationCountResponse{}

	
	counts, err := services.GetPublicationCount()
	if err != nil {
		if err == services.ErrPublicationCountFilterNotFound {
			return nil, huma.Error404NotFound("No publication counts found")
		}
		return nil, huma.Error500InternalServerError("Failed to retrieve publication counts")
	}

	
	response.Body = counts
	return response, nil
}