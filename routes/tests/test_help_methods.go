package tests

import (
	"errors"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"
	"log"
)

func GetBibliometrics(scientists responses.ScientistBody) *responses.BibliometricBody {

	result, err := services.GetBibliometricByScientistID(scientists.ID)
	if err != nil {
		log.Fatalf("GetBibliometrics error: %v", err)
	}
	if result == nil {
		log.Fatalf("GetBibliometrics result jest nil")

	}
	return result
}
func GetOrganization(scientist responses.ScientistBody) []responses.OrganizationBodyExtended {

	result, err := services.GetOrganizationsByScientistID(scientist.ID)
	if err != nil {
		log.Fatalf("GetOrganizationsByScientistID error: %v", err)
	}
	if result == nil {
		log.Fatalf("GetOrganizationsByScientistID result jest nil")
	}
	return result
}
func ContainsOrganization(organizations []responses.OrganizationBodyExtended, organizationName string) bool {
	for _, organization := range organizations {
		if organization.Name == organizationName {
			return true
		}
	}
	return false
}
func ContainsResearchArea(researchArea string, area []responses.ResearchArea) bool {
	for _, resArea := range area {
		if *resArea.Name == researchArea {
			return true
		}
	}
	return false
}
func ComparePublicationScores(expected []responses.PublicationScore, received []responses.PublicationScore) error {

	if len(received) != len(expected) {
		log.Printf("Oczekiwana długość %d, otrzymana długość %d", len(expected), len(received))
		return errors.New("niewłaściwa długość tablicy PublicationScore")
	}
	for i := 0; i < len(received); i++ {
		for j := 0; j < len(received); j++ {
			if *expected[i].Year == *received[j].Year {
				if *expected[i].Score != *received[j].Score {
					log.Printf("Oczekiwano punktacji %f z roku %d, a otrzymano %f z roku %d", *expected[i].Score, *expected[i].Year, *received[j].Score, *received[j].Year)
					return errors.New("nieprawidłowy wynik za rok")
				}
			}
		}

	}
	return nil
}
func CompareResearchAreas(expected []responses.ResearchArea, received []responses.ResearchArea) error {
	if len(received) != len(expected) {
		log.Printf("Oczekiwano %d, dyscyplin, a otrzymano %d", len(expected), len(received))
		return errors.New("nieodpowiednia ilość dyscyplin")
	}
	for i := 0; i < len(received); i++ {
		if *received[i].Name != *expected[i].Name {
			log.Printf("Oczekiwano dyscypliny %s, a otrzymano %s", *received[i].Name, *expected[i].Name)
			return errors.New("niepasujące nazwy dyscyplin")
		}
	}
	return nil
}
func ContainsPickedYearOfPublication(yearsOfPublication []responses.PublicationScore, expectedYear *int) error {
	for i := 0; i < len(yearsOfPublication); i++ {
		if yearsOfPublication[i].Year == expectedYear {
			return nil
		}
	}
	return errors.New("brak publikacji z oczekiwanego roku")
}
func ContainsPickedJournalType(publications []responses.PublicationBody, journalType string) error {
	for i := 0; i < len(publications); i++ {

		if *publications[i].Journal == journalType {
			return nil
		}
	}
	return errors.New("brak oczekiwanego typu publikacji")
}
