package handlers

import (
	"context"
	"io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetBibliometricByID(ctx context.Context, input *requests.BibliometricsID) (*responses.BibliometricsResponse, error) {
	logging.Logger.Info("INFO: Handling GetBibliometricByID request")
	response := &responses.BibliometricsResponse{}
	resultingBibliometrics, err := services.GetBibliometricByID(input.ID)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to get bibliometrics by ID:", err)
		return nil, huma.Error400BadRequest("Failed to get bibliometrics by ID")
	}
	logging.Logger.Info("INFO: Successfully retrieved bibliometrics by ID")
	response.Body = resultingBibliometrics
	return response, nil
}

func GetBibliometricByAuthor(ctx context.Context, input *requests.BibliometricsScientistID) (*responses.ListOfBibliometricsResponse, error) {
	logging.Logger.Info("INFO: Handling GetBibliometricByAuthor request")
	response := &responses.ListOfBibliometricsResponse{}
	resultingBibliometrics, err := services.GetBibliometricByAuthor(input.ID)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to get bibliometrics by author:", err)
		return nil, huma.Error400BadRequest("Failed to get bibliometrics by author")
	}
	logging.Logger.Info("INFO: Successfully retrieved bibliometrics by author")
	response.Body = resultingBibliometrics
	return response, nil
}

func CreateBibliometric(ctx context.Context, input *requests.CreateBibliometric) (*responses.CreateBibliometricResponse, error) {
	logging.Logger.Info("INFO: Handling CreateBibliometric request")
	response := &responses.CreateBibliometricResponse{}
	createdBibliometric, err := services.CreateBibliometric(input)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to create bibliometric:", err)
		return nil, huma.Error400BadRequest("Failed to create bibliometric")
	}
	logging.Logger.Info("INFO: Successfully created bibliometric")
	response.Body = responses.CreateBibliometric{ID: createdBibliometric}
	return response, nil
}

func DeleteBibliometricByID(ctx context.Context, input *requests.DeleteBibliometric) error {
	logging.Logger.Info("INFO: Handling DeleteBibliometricByID request")
	err := services.DeleteBibliometricByID(input)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to delete bibliometric by ID:", err)
		return err
	}
	logging.Logger.Info("INFO: Successfully deleted bibliometric by ID")
	return nil
}

func UpdateBibliometric(ctx context.Context, input *requests.UpdateBibliometric) (*responses.UpdateBibliometricResponse, error) {
	logging.Logger.Info("INFO: Handling UpdateBibliometric request")
	updatedBibliometric, err := services.UpdateBibliometricById(input)
	if err != nil {
		logging.Logger.Error("ERROR: Failed to update bibliometric:", err)
		return nil, huma.Error400BadRequest("Failed to update bibliometric")
	}
	logging.Logger.Info("INFO: Successfully updated bibliometric")
	return updatedBibliometric, nil
}
