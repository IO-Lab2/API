package internal

import (
	"context"
	"net/http"

	"io-project-api/internal/responses"
	"io-project-api/routes/bibliometrics"
	"io-project-api/routes/organizations"
	"io-project-api/routes/publications"
	"io-project-api/routes/scientists"
	scientistsorganizations "io-project-api/routes/scientists_organizations"
	scientistspublications "io-project-api/routes/scientists_publications"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterAPIRoutes(api huma.API, prefix string) {
	huma.Register(api, huma.Operation{
		OperationID: "Docs Redirect",
		Description: "Redirect to the API documentation",
		Tags:        []string{"Docs"},
		Method:      http.MethodGet,
		Path:        "/",

		Responses: map[string]*huma.Response{
			"302": {Description: "Redirect to docs"},
		},
		DefaultStatus: 302,
	}, func(ctx context.Context, input *struct{}) (*responses.RedirectResponse, error) {
		return &responses.RedirectResponse{Url: "/docs"}, nil
	})
	huma.Register(api, huma.Operation{
		OperationID: "Health Check",
		Description: "Check if the service is healthy",
		Tags:        []string{"Healthcheck"},
		Method:      http.MethodGet,
		Path:        "/healthcheck",
		Responses: map[string]*huma.Response{
			"200": {Description: "Healthy"},
		},
		DefaultStatus: 200,
	}, func(ctx context.Context, i *struct{}) (*responses.Healthcheck, error) {
		return &responses.Healthcheck{Message: "Healthy"}, nil
	})

	bibliometrics.RegisterBibliometricsRoutes(api, prefix)
	organizations.RegisterOrganizationsRoutes(api, prefix)
	publications.RegisterPublicationsRoutes(api, prefix)
	scientists.RegisterScientistsRoutes(api, prefix)
	scientistsorganizations.RegisterScientistsOrganizationsRoutes(api, prefix)
	scientistspublications.RegisterScientistsOrganizationsRoutes(api, prefix)
}
