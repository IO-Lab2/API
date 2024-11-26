package responses

import "io-project-api/internal/models"

type PublicationCountResponse struct {
	Body []models.PublicationCount `json:"body"`
}