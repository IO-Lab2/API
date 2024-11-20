package citationCount

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)


func RegisterCitationCountRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Citation Counts",
		Description: "Retrieve a list of citation counts",
		Tags:        []string{"citation-counts"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/citation-counts", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Citation counts retrieved successfully"},
			"404": {Description: "No citation counts found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.CitationCountFilterRequest) (*responses.CitationCountResponse, error) {
			return handlers.GetCitationCountHandler(ctx)
		},
	)
}