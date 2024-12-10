package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func GetPublicationByID(ctx context.Context, input *requests.PublicationID) (*responses.PublicationResponse, error) {
	logging.Logger.Info("INFO: Handling GetPublicationByID request")

	response := &responses.PublicationResponse{}
	result, err := services.GetPublicationByID(input.ID)
	if err != nil {
		if err == services.ErrPublicationNotFound {
			logging.Logger.Warn("WARNING: Publication not found", err)
			return nil, huma.NewError(http.StatusNotFound, "Publication not found")
		}

		logging.Logger.Error("ERROR: Failed to get publication by ID:", err)
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to get publication by ID")
	}

	logging.Logger.Info("INFO: Successfully retrieved publication by ID")
	response.Body = result
	return response, nil
}
func CreatePublication(ctx context.Context, input *requests.CreatePublicationRequest) (*responses.CreatePublicationResponse, error) {
	logging.Logger.Info("INFO: Handling CreatePublication request")

	response := &responses.CreatePublicationResponse{}
	id, err := services.CreatePublication(input)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to create publication:", err)
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to create publication")
	}

	logging.Logger.Info("INFO: Successfully created publication")
	response.Body = responses.CreatePublication{ID: id}
	return response, nil
}
func DeletePublication(ctx context.Context, input *requests.DeletePublication) error {
	logging.Logger.Info("INFO: handling DeletePublication request")

	err := services.DeletePublication(input)

	if err != nil {
		logging.Logger.Error("ERROR: Failed to delete publication")
	}
	logging.Logger.Info("Succesfully deleted publication")
	return nil
}
