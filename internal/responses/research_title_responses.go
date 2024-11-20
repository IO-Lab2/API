package responses

import "io-project-api/internal/models"

type ResearchTitleResponse struct {
	Body []models.ResearchTitle `json:"body"`
}