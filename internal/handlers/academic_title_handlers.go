package handlers

import (
	"context"
	"io-project-api/internal/logger"

	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetAcademicTitleHandler(ctx context.Context) (*responses.AcademicTitleResponse, error) {
	logging.Logger.Info("INFO: Handling GetAcademicTitle request")
	response := &responses.AcademicTitleResponse{}

	titles, err := services.GetAcademicTitles()
	if err != nil {
		if err == services.ErrAcademicTitleFilterNotFound {
			logging.Logger.Error("ERROR: No academic titles found")
			return nil, huma.Error404NotFound("No academic titles found")
		}
		logging.Logger.Error("ERROR: Failed to retrieve academic titles:", err)
		return nil, huma.Error500InternalServerError("Failed to retrieve academic titles")
	}
	logging.Logger.Info("INFO: Successfully retrieved academic titles")
	response.Body = titles
	return response, nil
}
