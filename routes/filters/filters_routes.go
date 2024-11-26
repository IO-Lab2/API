package filters

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterFiltersRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Publications Counts Filter",
		Description: "Retrieves a range of publication counts",
		Tags:        []string{"Filters", "Publication Counts"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/publication-counts", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Publication counts retrieved successfully"},
			"404": {Description: "No publication counts found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.PublicationCountFilterRequest) (*responses.PublicationCountResponse, error) {
			return handlers.GetPublicationCountHandler(ctx)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Ministerial Scores Filter",
		Description: "Retrieves the largest and smallest ministerial scores",
		Tags:        []string{"Filters", "Ministerial Scores"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/ministerial-scores", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Ministerial scores retrieved successfully"},
			"404": {Description: "No ministerial scores found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.MinisterialScoreFilterRequest) (*responses.MinisterialScoreResponse, error) {
			return handlers.GetMinisterialScoreHandler(ctx)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Organizations Filter",
		Description: "Retrieves the list of organizations",
		Tags:        []string{"Filters", "Organizations"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/organizations", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Organizations retrieved successfully"},
			"404": {Description: "No organizations found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.OrganizationFilterRequest) (*responses.ListOfOrganizationsResponse, error) {
			return handlers.GetOrganizationsHandler(ctx)

		},
	)
}
