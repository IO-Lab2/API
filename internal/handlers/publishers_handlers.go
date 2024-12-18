package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetPublishersHandler(ctx context.Context) (*responses.PublishersResponse, error) {
	logging.Logger.Info("INFO: Handling GetPublishers request")
	responses := &responses.PublishersResponse{}
	publishers, err := services.GetPublishers()
	if len(publishers) == 0 || err != nil {
		logging.Logger.Error("ERROR: Failed to retrieve publishers:", err)
		return nil, huma.Error400BadRequest("Failed to retrieve publishers")
	}

	logging.Logger.Info("INFO: Successfully retrieved publishers")

	responses.Body = publishers
	return responses, nil
}
