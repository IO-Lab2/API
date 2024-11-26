package responses

import "io-project-api/internal/models"

type MinisterialScoreResponse struct {
	Body []models.MinisterialScore `json:"body"`
}