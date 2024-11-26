package publicationCount

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterPublicationCountRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Publication Counts Filter",
		Description: "Retrieve a list of publication counts",
		Tags:        []string{"Publication Counts"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/publication-counts", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Publication counts retrieved successfully"},
			"404": {Description: "No publication counts found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.PublicationCountFilterRequest) (*responses.PublicationCountResponse, error) {
			return handlers.GetPublicationCountHandler(ctx)
		},
	)
}
