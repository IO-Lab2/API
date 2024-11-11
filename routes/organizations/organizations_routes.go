package organizations

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterOrganizationRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Organization by ID",
		Description: "Get Organization by ID",
		Tags:        []string{"Organization"},
		Path:        fmt.Sprintf("%s/bibliometrics/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Organization found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, input *requests.OrganizationID) (*responses.OrganizationsResponse, error) {
			return handlers.GetOrganizationById(ctx, input)
		})
}
