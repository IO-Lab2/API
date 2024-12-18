package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetPublicationsYearsHandler(ctx context.Context) (*responses.PublicationsYearsResponse, error) {
	logging.Logger.Info("INFO: Handling GetPublicationsYears request")
	response := &responses.PublicationsYearsResponse{}
	years, err := services.GetPublicationsYears()
	if len(years) == 0 || err != nil {
		logging.Logger.Error("ERROR: Failed to retrieve publications years:", err)
		return nil, huma.Error400BadRequest("Failed to retrieve publications years")
	}

	logging.Logger.Info("INFO: Successfully retrieved publications years")
	response.Body = years
	return response, nil
}
