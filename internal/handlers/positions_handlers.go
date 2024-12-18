package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetPositionsHandler(ctx context.Context) (*responses.PositionsResponse, error) {
	logging.Logger.Info("INFO: Handling GetPositions request")
	response := &responses.PositionsResponse{}
	positions, err := services.GetPositions()
	if len(positions) == 0 || err != nil {
		logging.Logger.Error("ERROR: Failed to retrieve positions:", err)
		return nil, huma.Error400BadRequest("Failed to retrieve positions")
	}

	logging.Logger.Info("INFO: Successfully retrieved positions")
	response.Body = positions
	return response, nil
}
