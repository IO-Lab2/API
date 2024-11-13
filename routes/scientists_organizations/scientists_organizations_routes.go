package scientistsorganizations

import (
	"context"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterScientistsOrganizationsRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get scientist organization by ID",
		Description: "Get a scientist organization by ID",
		Tags:        []string{"scientists organizations"},
		Method:      http.MethodGet,
		Responses: map[string]*huma.Response{
			"200": {Description: "Scientist organization found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, i *requests.ScientistOrganizationID) (*responses.ScientistOrganizationResponse, error) {
			return handlers.GetScientistOrganizationByID(ctx, i)
		})
}
