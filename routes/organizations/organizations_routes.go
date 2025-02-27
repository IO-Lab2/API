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
		Tags:        []string{"Organizations"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/organizations/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Organization found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, input *requests.OrganizationID) (*responses.OrganizationResponse, error) {
			return handlers.GetOrganizationById(ctx, input)
		})

	huma.Register(api, huma.Operation{
		OperationID: "Get Scientist Organization by ID",
		Description: "Get a scientist organization by ID",
		Tags:        []string{"Organizations"},
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
		Tags:        []string{"Organizations"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/organizations/scientist/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Organizations found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, input *requests.ScientistID) (*responses.ListOfOrganizationsResponse, error) {
			return handlers.GetOrganizationsByScientistId(ctx, input)
		})
	huma.Register(api, huma.Operation{
		OperationID: "Create Organization",
		Description: "Create Organization",
		Tags:        []string{"Organizations"},
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("%s/organizations", basePath),
		Responses: map[string]*huma.Response{
			"201": {Description: "Successfully created new organization"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, i *requests.CreateOrganization) (*responses.CreateOrganizationResponse, error) {
			return handlers.CreateOrganization(ctx, i)
		})
	huma.Register(api, huma.Operation{
		OperationID: "Update Organization",
		Description: "Update Organization",
		Tags:        []string{"Organizations"},
		Method:      http.MethodPut,
		Path:        fmt.Sprintf("%s/organizations/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Successfully updated organization"},
			"400": {Description: "Bad request"},
			"404": {Description: "Organization not found"},
		}},
		func(ctx context.Context, input *requests.UpdateOrganization) (*responses.UpdateOrganizationResponse, error) {
			return handlers.UpdateOrganization(ctx, input)
		})
	huma.Register(api, huma.Operation{
		OperationID: "Delete Organization",
		Description: "Delete Organization",
		Tags:        []string{"Organizations"},
		Method:      http.MethodDelete,
		Path:        fmt.Sprintf("%s/organizations/{id}", basePath),
		Responses: map[string]*huma.Response{
			"204": {Description: "Successfully deleted organization"},
			"400": {Description: "Bad request"},
			"404": {Description: "Organization not found"},
		}},
		func(ctx context.Context, input *requests.DeleteOrganization) (*huma.Response, error) {
			return &huma.Response{}, handlers.DeleteOrganization(ctx, input)
		})
}
