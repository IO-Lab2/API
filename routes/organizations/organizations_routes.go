package organizations

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterOrganizationsRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Organization by ID",
		Description: "Get Organization by ID",
		Tags:        []string{"Organization"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/organizations/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Organization found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, input *requests.OrganizationID) (*responses.OrganizationsResponse, error) {
			return handlers.GetOrganizationById(ctx, input)
		})

	huma.Register(api, huma.Operation{
		OperationID: "Get scientist organization by ID",
		Description: "Get a scientist organization by ID",
		Tags:        []string{"scientists organizations"},
		Method:      http.MethodGet,
		Path:        basePath + "/scientists_organizations/{id}",
		Responses: map[string]*huma.Response{
			"200": {Description: "Scientist organization found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, i *requests.ScientistOrganizationID) (*responses.ScientistOrganizationResponse, error) {
			return handlers.GetScientistOrganizationByID(ctx, i)
		})

	huma.Register(api, huma.Operation{
		OperationID: "Get Organizations By Scientist ID",
		Description: "Get Organizations By Scientist ID",
		Tags:        []string{"Organization"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/organizations/scientist/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Organizations found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, input *requests.ScientistID) (*responses.ListOfOrganizationsResponse, error) {
			return handlers.GetOrganizationsByScientistId(ctx, input)
		})
}
