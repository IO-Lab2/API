package ministerialscores

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterMinisterialScoresRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Ministerial Scores",
		Description: "Retrieve a list of ministerial scores",
		Tags:        []string{"ministerial-scores"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/ministerial-scores", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Ministerial scores retrieved successfully"},
			"404": {Description: "No ministerial scores found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.MinisterialScoreFilterRequest) (*responses.MinisterialScoresResponse, error) {
			return handlers.GetMinisterialScoresHandler(ctx)
		},
	)
}