package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetResearchTitleHandler(ctx context.Context) (*responses.ResearchAreaExtendedResponse, error) {
	logging.Logger.Info("INFO: Handling GetResearchTitleHandler request")
	response := &responses.ResearchAreaExtendedResponse{}

	areas, err := services.GetResearchAreas()
	if err != nil {
		if err == services.ErrResearchTitleFilterNotFound {
			logging.Logger.Error("ERROR: No research titles found")
			return nil, huma.Error404NotFound("No research titles found")
		}
		logging.Logger.Error("ERROR: Failed to retrieve research titles:", err)
		return nil, huma.Error500InternalServerError("Failed to retrieve research titles")
	}

	logging.Logger.Info("INFO: Successfully retrieved research titles")
	response.Body = areas
	return response, nil
}
