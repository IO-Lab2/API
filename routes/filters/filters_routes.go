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
		func(ctx context.Context, input *requests.OrganizationFilterRequest) (*responses.ListOfOrganizations, error) {
			return handlers.GetOrganizationsHandler(ctx)

		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Academic Titles Filter",
		Description: "Retrieves a list of academic titles",
		Tags:        []string{"Filters", "Academic Titles"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/academic-titles", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Academic titles retrieved successfully"},
			"404": {Description: "No academic titles found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.AcademicTitleFilterRequest) (*responses.AcademicTitleResponse, error) {
			return handlers.GetAcademicTitleHandler(ctx)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Citations Counts",
		Description: "Retrieves a range of citations counts",
		Tags:        []string{"Filters", "Citation Counts"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/citation-counts", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Citation counts retrieved successfully"},
			"404": {Description: "No citation counts found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.CitationCountFilterRequest) (*responses.CitationCountResponse, error) {
			return handlers.GetCitationCountHandler(ctx)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Research Titles Filter",
		Description: "Retrieves a list of research titles",
		Tags:        []string{"Filters", "Research Titles"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/research-titles", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Research titles retrieved successfully"},
			"404": {Description: "No research titles found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.ResearchAreasFilterRequest) (*responses.ResearchAreasResponse, error) {
			return handlers.GetResearchTitleHandler(ctx)
		},
	)
}
