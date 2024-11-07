package internal

import (
	"context"
	"net/http"

	"io-project-api/internal/responses"

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
}
