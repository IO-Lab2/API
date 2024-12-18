package responses

import (
	"io-project-api/internal/models"
)

// PublicationsYearsResponse represents the response for the publications years endpoint
type PublicationsYearsResponse struct {
	Body []models.PublicationsYear `json:"body"`
}
