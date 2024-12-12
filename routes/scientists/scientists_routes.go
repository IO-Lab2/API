package scientists

import (
	"context"
	"fmt"
	"net/http"

	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterScientistsRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Scientists by ID",
		Description: "Get a scientist by ID",
		Tags:        []string{"Scientists"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/scientists/{id}", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Scientist found"},
			"400": {Description: "Bad request"},
		},
	}, func(ctx context.Context, input *requests.ScientistID) (*responses.ScientistResponse, error) {

		return handlers.GetScientistByID(ctx, input)
	})
}
