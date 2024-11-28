package handlers

import (
	"context"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/zap"
)

func GetPublicationByID(ctx context.Context, input *requests.PublicationID) (*responses.PublicationResponse, error) {
	// Call the service layer to get the publication by ID
	response := &responses.PublicationResponse{}
	result, err := services.GetPublicationByID(input.ID)
	if err != nil {
		if err == services.ErrPublicationNotFound {
			zap.L().Warn("Publication not found", zap.Error(err))
			return nil, huma.NewError(http.StatusNotFound, "Publication not found")
		}

		zap.L().Error("Failed to get publication by ID", zap.Error(err))
		return nil, huma.NewError(http.StatusInternalServerError, "Failed to get publication by ID")
	}
	response.Body = result

	return response, nil
}
