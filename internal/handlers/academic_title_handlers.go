package handlers

import (
	"context"

	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetAcademicTitleHandler(ctx context.Context) (*responses.AcademicTitleResponse, error) {
	response := &responses.AcademicTitleResponse{}

	titles, err := services.GetAcademicTitles()
	if err != nil {
		if err == services.ErrAcademicTitleFilterNotFound {
			return nil, huma.Error404NotFound("No academic titles found")
		}
		return nil, huma.Error500InternalServerError("Failed to retrieve academic titles")
	}

	response.Body = titles
	return response, nil
}
