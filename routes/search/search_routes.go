package search

import (
	"context"
	"io-project-api/internal/handlers"
	"io-project-api/internal/models"
	"io-project-api/internal/responses"

	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterSearchRoutes(api huma.API, prefix string) {

	huma.Register(api, huma.Operation{
		OperationID: "Search",
		Description: "Search for academic profiles.",
		Tags:        []string{"Search"},
		Method:      http.MethodGet,
		Path:        prefix + "/search",
		Responses: map[string]*huma.Response{
			"200": {Description: "Search results retrieved successfully"},
			"400": {Description: "Invalid search input"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *models.SearchInput) (*responses.ScientistsResponse, error) {

			if input.MinImpactFactor < 0 || input.MinImpactFactor > input.MaxImpactFactor {
				return nil, huma.Error400BadRequest("Invalid search input")
			}

			if input.MinMinisterialScore < 0 || input.MinMinisterialScore > input.MaxMinisterialScore {
				return nil, huma.Error400BadRequest("Invalid search input")
			}

			if input.MinPublications < 0 || input.MinPublications > input.MaxPublications {
				return nil, huma.Error400BadRequest("Invalid search input")
			}

			return handlers.SearchHandler(ctx, input)
		},
	)
}
