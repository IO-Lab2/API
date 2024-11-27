package responses

import "io-project-api/internal/models"

type ResearchAreasResponse struct {
	Body []models.ResearchArea `json:"body" doc:"Research areas"`
}
