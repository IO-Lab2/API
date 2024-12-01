package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetPublicationCountHandler(ctx context.Context) (*responses.PublicationCountResponse, error) {
	logging.Logger.Info("INFO: Handling GetPublicationCount request")
	response := &responses.PublicationCountResponse{}
	counts, err := services.GetPublicationCount()
	if err != nil {
		if err == services.ErrPublicationCountFilterNotFound {
			logging.Logger.Error("ERROR: No publication counts found")
			return nil, huma.Error404NotFound("No publication counts found")
		}
		logging.Logger.Error("ERROR: Failed to retrieve publication counts:", err)
		return nil, huma.Error500InternalServerError("Failed to retrieve publication counts")
	}

	logging.Logger.Info("INFO: Successfully retrieved publication counts")
	response.Body = counts
	return response, nil
}
