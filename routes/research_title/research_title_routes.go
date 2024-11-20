package researchTitle

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)


func RegisterResearchTitleRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Research Titles",
		Description: "Retrieve a list of research titles",
		Tags:        []string{"research-titles"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/research-titles", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Research titles retrieved successfully"},
			"404": {Description: "No research titles found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.ResearchTitleFilterRequest) (*responses.ResearchTitleResponse, error) {
			return handlers.GetResearchTitleHandler(ctx)
		},
	)
}