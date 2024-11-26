package responses

import "io-project-api/internal/models"

type AcademicTitleResponse struct {
	Body []models.AcademicTitle `json:"body" doc:"Academic titles"`
}
