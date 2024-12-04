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
			return handlers.SearchHandler(ctx, input)
		},
	)
}
