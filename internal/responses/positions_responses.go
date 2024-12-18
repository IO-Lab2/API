package responses

import "io-project-api/internal/models"

type PositionsResponse struct {
	Body []models.PositionFilter `json:"body" doc:"List of positions"`
}
