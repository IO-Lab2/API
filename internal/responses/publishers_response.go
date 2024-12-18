package responses

import "io-project-api/internal/models"

type PublishersResponse struct {
	Body []models.PublisherFilter `json:"body" doc:"List of publishers"`
}
