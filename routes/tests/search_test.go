package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io-project-api/internal/models"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearchAcademicTitle(t *testing.T) {

	router := TestSetUP()

	url := "http://localhost:8000/api/search?academic_titles%5B%5D=PhD"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}

	for _, item := range result {
		if err := item.AcademicTitle == "PhD"; err == false {
			t.Errorf("AcademicTitle jest niewłaściwy.")
		}
	}
}

func TestSearchMinisterialScore(t *testing.T) {

	router := TestSetUP()

	minimalScore := 15.5
	maximalScore := 205.5

	url := fmt.Sprintf("http://localhost:8000/api/search?ministerial_score_min=%v&ministerial_score_max=%v", minimalScore, maximalScore)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}

	for _, item := range result {
		bibliometric := GetBibliometrics(item)
		if bibliometric.MinisterialScore < minimalScore || bibliometric.MinisterialScore > maximalScore {
			t.Errorf("Niewłaściwa punktacja")
		}
	}
}
func GetBibliometrics(scientists models.Scientist) *responses.BibliometricBody {

	result, err := services.GetBibliometricByScientistID(scientists.ID)
	if err != nil {
		log.Fatalf("GetBibliometrics error: %v", err)
	}
	if result == nil {
		log.Fatalf("GetBibliometrics result jest nil")

	}
	return result
}
func TestSearchByName(t *testing.T) {

	router := TestSetUP()

	name := "Marcin"
	url := fmt.Sprintf("http://localhost:8000/api/search?name=%s", name)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result {
		if item.FirstName != name {
			t.Errorf("Imię naukowca się niezgadza.")
		}
	}

}
func TestSearchByOrganizations(t *testing.T) {
	router := TestSetUP()
	organizationName := "Institute of Information Technology" //Nazwa potrzebna do stworzenia zapytania

	url := "http://localhost:8000/api/search?organizations%5B%5D=Institute of Information Technology"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result {
		organizations := GetOrganization(item)
		if ContainsOrganization(organizations, organizationName) == false {
			t.Errorf("Niewłaściwa organizacja naukowca.")
		}

	}
}
func GetOrganization(scientist models.Scientist) []responses.OrganizationBodyExtended {

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

func TestSearchByJournalTypes(t *testing.T) {
	t.Skip("Test nieskończony. Czeka na wprowadzenie funkcjonalności.")
	router := TestSetUP()
	journalType := "artykuł"

	url := fmt.Sprintf("http://localhost:8000/api/search?journal_type=%s", journalType)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result {
		bibliometric := GetBibliometrics(item)
		if bibliometric == bibliometric {
			//dokończyć
		}
	}
}
func TestSearchByPositions(t *testing.T) {

	router := TestSetUP()

	position := "Professor"
	url := "http://localhost:8000/api/search?positions%5B%5D=" + position
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result {
		if item.Position == nil {
			log.Fatalf("Przekazano pusy wskaźnik.")
		}
		if position != *item.Position {
			t.Errorf("Niewłaściwa stanowisko naukowca.")
			t.SkipNow()
		}
	}

}
func TestSearchByPublicationsCount(t *testing.T) {
	router := TestSetUP()

	publicationMinCount := 5
	publicationMaxCount := 20

	url := fmt.Sprintf("http://localhost:8000/api/search?publications_min=%v&publications_max=%v", publicationMinCount, publicationMaxCount)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result {
		bibliometric := GetBibliometrics(item)
		if bibliometric.PublicationCount < publicationMinCount || bibliometric.PublicationCount > publicationMaxCount {
			t.Errorf("Niewłaściwa ilość publikacji")
		}
	}
}

func TestSearchByResearchAreas(t *testing.T) {
	router := TestSetUP()
	researchArea := "information and communication technology (ICT)"

	url := "http://localhost:8000/api/search?research_areas%5B%5D=" + researchArea
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result {
		if !ContainsResearchArea(researchArea, item.ResearchAreas) {
			t.Errorf("Naukowiec nie posiada odpowiedniej dyscypliny.")
			t.SkipNow()
		}
	}
}
func ContainsResearchArea(researchArea string, area []models.ResearchArea) bool {
	for _, resArea := range area {
		if resArea.Name == researchArea {
			return true
		}
	}
	return false
}
func TestSearchBySurname(t *testing.T) {
	router := TestSetUP()
	surname := "Kowa"

	url := fmt.Sprintf("http://localhost:8000/api/search?surname=%s", surname)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result {
		if !strings.Contains(strings.ToLower(item.LastName), strings.ToLower(surname)) {
			t.Errorf("Nazwisko naukowca nie zawiera fragmentu: %s.", surname)
			t.SkipNow()
		}
	}
	t.Logf("Test przeszedł pomyślnie")
}
func TestSearchByYearScoreFilters(t *testing.T) {
	t.Skip("Test nieskończony. Czeka na zaimplementowanie funkcjonalności.")
}
