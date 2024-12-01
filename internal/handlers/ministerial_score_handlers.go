package handlers

import (
	"context"
	"io-project-api/internal/logger"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetMinisterialScoreHandler(ctx context.Context) (*responses.MinisterialScoreResponse, error) {
	logging.Logger.Info("INFO: Handling GetMinisterialScore request")
	response := &responses.MinisterialScoreResponse{}

	scores, err := services.GetMinisterialScores()
	if err != nil {
		if err == services.ErrMinisterialScoreFilterNotFound {
			logging.Logger.Error("ERROR: No ministerial scores found")
			return nil, huma.Error404NotFound("No ministerial scores found")
		}
		logging.Logger.Error("ERROR: Failed to retrieve ministerial scores:", err)
		return nil, huma.Error500InternalServerError("Failed to retrieve ministerial scores")
	}

	logging.Logger.Info("INFO: Successfully retrieved ministerial scores")
	response.Body = scores
	return response, nil
}
