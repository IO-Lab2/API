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
		OperationID: "Get Research Areas Filter",
		Description: "Retrieves a list of research Areas",
		Tags:        []string{"Filters", "Research Areas"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/research-areas", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Research areas retrieved successfully"},
			"404": {Description: "No research areas found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.ResearchAreasFilterRequest) (*responses.ResearchAreaExtendedResponse, error) {

			return handlers.GetResearchTitleHandler(ctx)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Traverse Organizations Tree",
		Description: "Traverses the organizations tree for a nice visualization of the organizations",
		Tags:        []string{"Filters", "Organizations Tree"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/organizations-tree", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Organizations tree retrieved successfully"},
			"404": {Description: "No organizations tree found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.OrganizationTreeFilterRequest) (*responses.ListOfOrganizations, error) {

			return handlers.GetOrganizationTreeHandler(ctx, input)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Positions Filter",
		Description: "Retrieves a list of positions",
		Tags:        []string{"Filters", "Positions"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/positions", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Positions retrieved successfully"},
			"404": {Description: "No positions found"},
			"500": {Description: "Internal server error"},
		}},

		func(ctx context.Context, input *requests.PositionFilterRequest) (*responses.PositionsResponse, error) {
			return handlers.GetPositionsHandler(ctx)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Publishers Filter",
		Description: "Retrieves a list of publishers",
		Tags:        []string{"Filters", "Publishers"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/publishers", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Publishers retrieved successfully"},
			"404": {Description: "No publishers found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.PublishersFilterRequest) (*responses.PublishersResponse, error) {
			return handlers.GetPublishersHandler(ctx)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Publication Years Filter",
		Description: "Retrieves a list of publication years",
		Tags:        []string{"Filters", "Publication Years"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/publication-years", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Publication years retrieved successfully"},
			"404": {Description: "No publication years found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.PublicationsYearsFilterRequest) (*responses.PublicationsYearsResponse, error) {
			return handlers.GetPublicationsYearsHandler(ctx)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Journal Types Filter",
		Description: "Retrieves a list of journal types",
		Tags:        []string{"Filters", "Journal Types"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/journal-types", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Journal types retrieved successfully"},
			"404": {Description: "No journal types found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.JournalTypesFilterRequest) (*responses.JournalTypeResponse, error) {
			return handlers.GetJournalTypesHandler(ctx)
		},
	)

	huma.Register(api, huma.Operation{
		OperationID: "Get Impact Factors Filter",
		Description: "Retrieves the largest and smallest impact factors",
		Tags:        []string{"Filters", "Impact Factors"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/filters/impact-factors", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Impact factors retrieved successfully"},
			"404": {Description: "No impact factors found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.ImpactFactorsFilterRequest) (*responses.ImpactFactorResponse, error) {
			return handlers.GetImpactFactorsHandler(ctx)
		},
	)
}
