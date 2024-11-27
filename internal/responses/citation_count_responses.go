package responses

import "io-project-api/internal/models"

type CitationCountResponse struct {
	Body *models.CitationsFilter `json:"body"`
}
