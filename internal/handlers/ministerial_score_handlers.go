package handlers

import (
	"context"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetMinisterialScoreHandler(ctx context.Context) (*responses.MinisterialScoreResponse, error) {
	response := &responses.MinisterialScoreResponse{}

	scores, err := services.GetMinisterialScores()
	if err != nil {
		if err == services.ErrMinisterialScoreFilterNotFound {
			return nil, huma.Error404NotFound("No ministerial scores found")
		}
		return nil, huma.Error500InternalServerError("Failed to retrieve ministerial scores")
	}

	response.Body = scores
	return response, nil
}
