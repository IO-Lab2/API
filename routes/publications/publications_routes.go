package publications

import (
	"context"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterPublicationsRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get publications by ID",
		Description: "Get a publication by ID",
		Tags:        []string{"publications"},
		Method:      http.MethodGet,
		Responses: map[string]*huma.Response{
			"200": {Description: "Publication found"},
			"400": {Description: "Bad request"},
		},
	}, func(ctx context.Context, i *requests.PublicationID) (*responses.PublicationsResponse, error) {
		return handlers.GetPublicationByID(ctx, i)
	})
}
