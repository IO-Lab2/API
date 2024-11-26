package acadmiecTitle

import (
	"context"
	"fmt"
	"io-project-api/internal/handlers"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterAcadmiecTitleRoutes(api huma.API, basePath string) {
	huma.Register(api, huma.Operation{
		OperationID: "Get Academic Titles",
		Description: "Retrieve a list of academic titles",
		Tags:        []string{"Academic Titles"},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("%s/academic-titles", basePath),
		Responses: map[string]*huma.Response{
			"200": {Description: "Academic titles retrieved successfully"},
			"404": {Description: "No academic titles found"},
			"500": {Description: "Internal server error"},
		}},
		func(ctx context.Context, input *requests.AcademicTitleFilterRequest) (*responses.AcademicTitleResponse, error) {
			return handlers.GetAcademicTitleHandler(ctx)
		},
	)
}
