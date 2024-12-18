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
		Tags:        []string{"Bibliometrics"},
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
		OperationID: "Get Bibliometrics by Scientist ID",
		Description: "Get Bibliometrics by Scientist ID",
		Tags:        []string{"Bibliometrics"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/bibliometrics/author/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Bibliometrics found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, input *requests.BibliometricsScientistID) (*responses.BibliometricsResponse, error) {

			return handlers.GetBibliometricByScientistID(ctx, input)
		},
	)
	huma.Register(api, huma.Operation{
		OperationID: "Create Bibliometrics",
		Description: "Create Biblometrics",
		Tags:        []string{"Bibliometrics"},
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("%s/bibliometrics", basePath),
		Responses: map[string]*huma.Response{
			"201": {Description: "Succesfully created new biblometric"},
			"400": {Description: "Failed to create bibliometric!"},
		},
	},
		func(ctx context.Context, input *requests.CreateBibliometric) (*responses.CreateBibliometricResponse, error) {
			return handlers.CreateBibliometric(ctx, input)
		})
	huma.Register(api, huma.Operation{
		OperationID: "Delete Bibliometrics",
		Description: "Delete Bibliometrics",
		Tags:        []string{"Bibliometrics"},
		Method:      http.MethodDelete,
		Path:        fmt.Sprintf("%s/bibliometrics/{id}", basePath),
		Responses: map[string]*huma.Response{
			"204": {Description: "Bibliometric deleted!"},
			"400": {Description: "Bad request"},
			"404": {Description: "Bibliometric not found"},
		},
	},
		func(ctx context.Context, input *requests.DeleteBibliometric) (*huma.Response, error) {
			return &huma.Response{}, handlers.DeleteBibliometricByID(ctx, input)
		})
	huma.Register(api, huma.Operation{
		OperationID: "Update Bibliometrics",
		Description: "Update bibliometrics",
		Tags:        []string{"Bibliometrics"},
		Method:      http.MethodPut,
		Path:        fmt.Sprintf("%s/bibliometrics/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Bibliometrics updated"},
			"400": {Description: "Bad request"},
			"404": {Description: "Bibliometric not found"},
		}},
		func(ctx context.Context, input *requests.UpdateBibliometric) (*responses.UpdateBibliometricResponse, error) {
			return handlers.UpdateBibliometric(ctx, input)
		},
	)
}
