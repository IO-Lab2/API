package services

import (
	"errors"
	"io-project-api/internal/database"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/repositories"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"

	"github.com/google/uuid"
)

var (
	ErrBibliometricNotFound = errors.New("bibliometric not found for the given ID")
)

func GetBibliometricByID(id uuid.UUID) (*responses.BibliometricBody, error) {
	logging.Logger.Info("INFO: Retrieving Bibliometric by ID")
	bibliometric, err := repositories.BibliometricByID(database.GetDB(), id)
	if err != nil {
		logging.Logger.Error("ERROR: Error querying bibliometric by ID", err)
		return nil, err
	}
	logging.Logger.Info("INFO: Successfully retrieving Bibliometric by ID")
	return bibliometric, nil
}

func GetBibliometricByScientistID(id uuid.UUID) (*responses.BibliometricBody, error) {
	logging.Logger.Info("INFO: Retrieving Bibliometric by Author")
	result, err := repositories.BibliometricByScientistID(database.GetDB(), id)
	if err != nil {
		logging.Logger.Error("ERROR: Error querying bibliometric by Author", err)
		return nil, err
	}

	logging.Logger.Info("INFO: Successfully retrieving Bibliometric by Author")
	return result, nil
}
func CreateBibliometric(input *requests.CreateBibliometric) (uuid.UUID, error) {
	id := uuid.New()
	logging.Logger.Info("INFO: Creating Bibliometric")
	err := repositories.CreateBibliometric(database.GetDB(), id, input)
	if err != nil {
		logging.Logger.Error("ERROR: Error creating bibliometric", err)
		return uuid.Nil, err
	}
	logging.Logger.Info("INFO: Successfully creating Bibliometric")
	return id, nil
}
func DeleteBibliometricByID(input *requests.DeleteBibliometric) error {
	logging.Logger.Info("INFO: Deleting Bibliometric by ID")
	err := repositories.DeleteBibliometric(database.GetDB(), input)
	if err != nil {
		logging.Logger.Error("ERROR: Error deleting Bibliometrics!", err)
		return err
	}
	logging.Logger.Info("INFO: Successfully deleting Bibliometric by ID")

	return nil
}
func UpdateBibliometricById(input *requests.UpdateBibliometric) (*responses.UpdateBibliometricResponse, error) {
	logging.Logger.Info("INFO: Updating Bibliometric by ID")
	err := repositories.UpdateBibliometric(database.GetDB(), input)
	if err != nil {
		logging.Logger.Error("ERROR: Error updating bibliometric!", err)
		return nil, err
	}
	updatedBibliometric, err := repositories.BibliometricByID(database.GetDB(), input.ID)
	if err != nil {
		logging.Logger.Error("ERROR: Error geting updated bibliometrics!", err)
		return nil, err
	}
	logging.Logger.Info("INFO: uccessfully updating Bibliometric by ID")

	return &responses.UpdateBibliometricResponse{Body: *updatedBibliometric}, nil
}
