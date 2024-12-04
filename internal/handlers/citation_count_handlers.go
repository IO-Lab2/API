package handlers

import (
	"context"
	"io-project-api/internal/logger"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetCitationCountHandler(ctx context.Context) (*responses.CitationCountResponse, error) {
	logging.Logger.Info("INFO: Handling GetCitationCount request")
	response := &responses.CitationCountResponse{}

	counts, err := services.GetCitationCount()
	if err != nil {
		if err == services.ErrCitationCountFilterNotFound {
			logging.Logger.Error("ERROR: No citation counts found")
			return nil, huma.Error404NotFound("No citation counts found")
		}
		logging.Logger.Error("ERROR: Failed to retrieve citation counts:", err)
		return nil, huma.Error500InternalServerError("Failed to retrieve citation counts")
	}

	logging.Logger.Info("INFO: Successfully retrieved citation counts")
	response.Body = counts
	return response, nil
}
