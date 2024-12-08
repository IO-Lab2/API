package publications

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterPublicationsRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Publications By ID",
		Description: "Get a publication by ID",
		Tags:        []string{"Publications"},
		Method:      http.MethodGet,
		Path:        basePath + "/publications/{id}",
		Responses: map[string]*huma.Response{
			"200": {Description: "Publication found"},
			"400": {Description: "Bad request"},
		},
	}, func(ctx context.Context, i *requests.PublicationID) (*responses.PublicationResponse, error) {
		return handlers.GetPublicationByID(ctx, i)
	})

	huma.Register(api, huma.Operation{
		OperationID: "Get Scientists Publication By Scientists ID",
		Description: "Get Scientists Publication By Scientists ID",
		Tags:        []string{"Publications"},
		Method:      http.MethodGet,
		Path:        basePath + "/scientists_publications/{id}",
		Responses: map[string]*huma.Response{
			"200": {Description: "Scientist publication found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, i *requests.ScientistPublicationID) (*responses.ScientistPublicationResponse, error) {
			return handlers.GetScientistPublicationByID(ctx, i)
		})
	huma.Register(api, huma.Operation{
		OperationID: "Create Publication",
		Description: "Create a publication",
		Tags:        []string{"Publications"},
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("%s/publications", basePath),
		Responses: map[string]*huma.Response{
			"201": {Description: "Successfully created new publication"},
			"400": {Description: "Failed to create publication!"},
		},
	},
		func(ctx context.Context, i *requests.CreatePublicationRequest) (*responses.CreatePublicationResponse, error) {
			return handlers.CreatePublication(ctx, i)
		},
	)
}
