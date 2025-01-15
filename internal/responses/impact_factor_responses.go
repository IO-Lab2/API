package responses

import "io-project-api/internal/models"

// ImpactFactorResponse defines the response body for the GetImpactFactorByID handler.
type ImpactFactorResponse struct {
	Body *ImpactFactorResponseBody `json:"body"`
}

// ImpactFactorResponseBody defines the response body for the GetImpactFactorByID handler.
type ImpactFactorResponseBody = models.RangeFilter
