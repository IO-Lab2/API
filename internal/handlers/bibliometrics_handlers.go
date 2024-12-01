package handlers

import (
	"context"
	logging "io-project-api/internal/logger"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetBibliometricByID(ctx context.Context, input *requests.BibliometricsID) (*responses.BibliometricsResponse, error) {
	response := &responses.BibliometricsResponse{}
	resultingBibliometrics, err := services.GetBibliometricByID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to get bibliometrics by ID")
	}
	response.Body = resultingBibliometrics
	return response, nil
}

func GetBibliometricByAuthor(ctx context.Context, input *requests.BibliometricsScientistID) (*responses.ListOfBibliometricsResponse, error) {
	response := &responses.ListOfBibliometricsResponse{}
	resultingBibliometrics, err := services.GetBibliometricByAuthor(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to get bibliometrics by author")
	}
	response.Body = resultingBibliometrics
	return response, nil
}
func CreateBibliometric(ctx context.Context, input *requests.CreateBibliometric) (*responses.CreateBibliometricResponse, error) {
	response := &responses.CreateBibliometricResponse{}
	createdBibliometric, err := services.CreateBibliometric(input)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to create bibliometric")
	}
	response.Body = responses.CreateBibliometric{ID: createdBibliometric}
	return response, nil
}
func DeleteBibliometricByID(ctx context.Context, input *requests.DeleteBibliometric) error {
	err := services.DeleteBibliometricByID(input)
	return err
}
func UpdateBibliometric(ctx context.Context, input *requests.UpdateBibliometric) (*responses.UpdateBibliometricResponse, error) {
	updatedBibliometric, err := services.UpdateBibliometricById(input)
	if err != nil {
		logging.Logger.Error("Failed to update bibliometrics", err)
		return nil, err
	}
	return updatedBibliometric, nil
}
