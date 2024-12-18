package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetJournalTypesHandler(ctx context.Context) (*responses.JournalTypeResponse, error) {
	logging.Logger.Info("INFO: Handling GetJournalTypes request")
	response := &responses.JournalTypeResponse{}
	journalTypes, err := services.GetJournalTypes()
	if err != nil {
		logging.Logger.Error("ERROR: Failed to retrieve journal types:", err)
		return nil, huma.Error500InternalServerError("Failed to retrieve journal types")
	}

	logging.Logger.Info("INFO: Successfully retrieved journal types")
	response.Body = journalTypes
	return response, nil
}
