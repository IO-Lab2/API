package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/models"
	"io-project-api/internal/repositories"
)

var ErrImpactFactorFilterNotFound = errors.New("no impact factors found")

// GetImpactFactors retrieves the impact factors from the database.
func GetImpactFactors() (*models.RangeFilter, error) {
	logging.Logger.Info("INFO: Retrieving ministerial counts")
	db := database.GetDB()

	scores, err := repositories.ImpactFactorFilter(db)
	if err != nil {
		logging.Logger.Error("ERROR: Error retrieving impact factors ", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieving impact factors")

	return scores, nil
}
