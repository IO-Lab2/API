package handlers

import (
	"context"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetResearchTitleHandler(ctx context.Context) (*responses.ResearchAreasResponse, error) {
	response := &responses.ResearchAreasResponse{}

	areas, err := services.GetResearchAreas()
	if err != nil {
		if err == services.ErrResearchTitleFilterNotFound {
			return nil, huma.Error404NotFound("No research titles found")
		}
		return nil, huma.Error500InternalServerError("Failed to retrieve research titles")
	}

	response.Body = areas
	return response, nil
}
