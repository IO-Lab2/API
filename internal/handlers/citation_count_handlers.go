package handlers

import (
	"context"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetCitationCountHandler(ctx context.Context) (*responses.CitationCountResponse, error) {
	response := &responses.CitationCountResponse{}

	counts, err := services.GetCitationCount()
	if err != nil {
		if err == services.ErrCitationCountFilterNotFound {
			return nil, huma.Error404NotFound("No citation counts found")
		}
		return nil, huma.Error500InternalServerError("Failed to retrieve citation counts")
	}

	response.Body = counts
	return response, nil
}
