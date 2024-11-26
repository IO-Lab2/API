package bibliometrics

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterBibliometricsRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Bibliometrics by ID",
		Description: "Get Bibliometrics by ID",
		Tags:        []string{"bibliometrics"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/bibliometrics/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Bibliometrics found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, input *requests.BibliometricsID) (*responses.BibliometricsResponse, error) {
			return handlers.GetBibliometricByID(ctx, input)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Bibliometrics by Author",
		Description: "Get Bibliometrics by Author",
		Tags:        []string{"bibliometrics"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/bibliometrics/author/{author}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Bibliometrics found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, input *requests.BibliometricsAuthor) (*responses.ListOfBibliometricsResponse, error) {

			return handlers.GetBibliometricByAuthor(ctx, input)
		},
	)
}
