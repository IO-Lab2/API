package scientistspublications

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
		OperationID: "Get scientist publication by ID",
		Description: "Get a scientist publication by ID",
		Tags:        []string{"scientists publications"},
		Method:      http.MethodGet,
		Path:        basePath + "/scientists_publications/{id}",
		Responses: map[string]*huma.Response{
			"200": {Description: "Scientist publication found"},
			"400": {Description: "Bad request"},
		}},
		func(ctx context.Context, i *requests.ScientistPublicationID) (*responses.ScientistPublicationResponse, error) {
			return handlers.GetScientistPublicationByID(ctx, i)
		})
}
