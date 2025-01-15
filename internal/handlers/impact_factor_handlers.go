package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetImpactFactorsHandler(ctx context.Context) (*responses.ImpactFactorResponse, error) {
	logging.Logger.Info("INFO: Handling GetImpactFactors request")
	response := &responses.ImpactFactorResponse{}

	factors, err := services.GetImpactFactors()
	if err != nil {
		if err == services.ErrImpactFactorFilterNotFound {
			logging.Logger.Error("ERROR: No impact factors found")
			return nil, huma.Error404NotFound("No impact factors found")
		}
		logging.Logger.Error("ERROR: Failed to retrieve impact factors:", err)
		return nil, huma.Error500InternalServerError("Failed to retrieve impact factors")
	}

	logging.Logger.Info("INFO: Successfully retrieved impact factors")
	response.Body = factors
	return response, nil
}
