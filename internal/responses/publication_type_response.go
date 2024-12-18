package responses

import "io-project-api/internal/models"

type JournalTypeResponse struct {
	Body []models.JournalType `json:"body"`
}
