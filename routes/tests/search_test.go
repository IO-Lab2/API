package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io-project-api/internal/database"
	"io-project-api/internal/repositories"
	"io-project-api/internal/responses"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearchAcademicTitle(t *testing.T) {

	router := TestSetUP()

	url := "http://localhost:8000/api/search?academic_titles%5B%5D=PhD"

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}

	for _, item := range result.Scientists {
		if err := *item.AcademicTitle == "PhD"; err == false {
			t.Errorf("AcademicTitle jest niewłaściwy.")
		}
	}
}
func TestSearchBadAcademicTitle(t *testing.T) {
	router := TestSetUP()
	url := "http://localhost:8000/api/search?academic_titles%5B%5D=BAD_TITLE"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwano kod błędu: %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}

func TestSearchMinisterialScore(t *testing.T) {

	router := TestSetUP()

	minimalScore := 15.5
	maximalScore := 205.5

	url := fmt.Sprintf("http://localhost:8000/api/search?ministerial_score_min=%f&ministerial_score_max=%f", minimalScore, maximalScore)

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}

	for _, item := range result.Scientists {
		bibliometric := GetBibliometrics(item)
		if *bibliometric.MinisterialScore < minimalScore || *bibliometric.MinisterialScore > maximalScore {
			t.Errorf("Niewłaściwa punktacja")
		}
	}
}

func TestSearchNegativeMinMinisterialScore(t *testing.T) {
	router := TestSetUP()
	minimalScore := -15.5
	maximalScore := 0

	url := fmt.Sprintf("http://localhost:8000/api/search?ministerial_score_min=%v&ministerial_score_max=%v", minimalScore, maximalScore)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwano kod błędu: %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}

func TestSearchNegativeMaxMinisterialScore(t *testing.T) {
	router := TestSetUP()
	minimalScore := 0
	maximalScore := -205.5

	url := fmt.Sprintf("http://localhost:8000/api/search?ministerial_score_min=%v&ministerial_score_max=%v", minimalScore, maximalScore)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwano kod błędu: %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}

func TestSearchMinisterialScoreWhenNotANumber(t *testing.T) {
	router := TestSetUP()
	minimalScore := "Pomidor"
	maximalScore := 205.5

	url := fmt.Sprintf("http://localhost:8000/api/search?ministerial_score_min=%v&ministerial_score_max=%v", minimalScore, maximalScore)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Oczekiwano kod błędu: %d, a otrzymano %d", http.StatusUnprocessableEntity, w.Code)
	}
}

func TestSearchByName(t *testing.T) {

	router := TestSetUP()

	name := "Marcin"
	url := fmt.Sprintf("http://localhost:8000/api/search?name=%s", name)

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result.Scientists {
		if *item.FirstName != name {
			t.Errorf("Imię naukowca się niezgadza.")
		}
	}

}

func TestSearchByBadName(t *testing.T) {
	router := TestSetUP()
	name := "WRONG_NAME"

	url := fmt.Sprintf("http://localhost:8000/api/search?name=%s", name)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania %d", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwano kodu błędu: %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}

func TestSearchByOrganizations(t *testing.T) {
	router := TestSetUP()
	organizationName := "Institute of Information Technology" //Nazwa potrzebna do stworzenia zapytania

	url := "http://localhost:8000/api/search?organizations%5B%5D=Institute of Information Technology"
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result.Scientists {
		organizations := GetOrganization(item)
		if ContainsOrganization(organizations, organizationName) == false {
			t.Errorf("Niewłaściwa organizacja naukowca dla %s %s o ID: %s", *item.FirstName, *item.LastName, item.ID)
		}

	}
}

func TestSearchByOrganizationsBadName(t *testing.T) {
	router := TestSetUP()
	organizationName := "WRONG ORGANIZATION NAME"

	url := "http://localhost:8000/api/search?organizations%5B%5D=" + organizationName
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żdania: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwano kodu błędu: %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}

func TestSearchByJournalTypes(t *testing.T) {

	router := TestSetUP()
	url := "http://localhost:8000/api/search?journal_types%5B%5D=artyku%C5%82"

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}

	db, err := database.InitDB()
	if err != nil {
		t.Errorf("Nie udało się połączyć z bazą danych")
	}
	journalType := "artykuł"
	for _, item := range result.Scientists {
		t.Logf("Niewłaściwe typy publikci naukowca dla %s %s o ID: %s", *item.FirstName, *item.LastName, item.ID)
		publications, err := repositories.PublicationsByScientistID(db, item.ID)
		if err != nil {
			t.Errorf("Nie udało się dostać publikacji: %d", err)
		}
		if ContainsPickedJournalType(publications, journalType) != nil {
			t.Errorf(err.Error())
		}

	}
}

func TestSearchByBadJournalTypes(t *testing.T) {
	router := TestSetUP()

	url := "http://localhost:8000/api/search?journal_types%5B%5D=BAD JOURNAL TYPE%C5%82"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwano kod błędu %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}

func TestSearchByPositions(t *testing.T) {

	router := TestSetUP()

	position := "Professor"
	url := "http://localhost:8000/api/search?positions%5B%5D=" + position
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result.Scientists {
		if item.Position == nil {
			log.Fatalf("Przekazano pusy wskaźnik.")
		}
		if position != *item.Position {
			t.Errorf("Niewłaściwa stanowisko naukowca.")
			t.SkipNow()
		}
	}

}

func TestSearchByBadPositions(t *testing.T) {
	router := TestSetUP()
	position := "BAD POSITION"

	url := "http://localhost:8000/api/search?positions%5B%5D=" + position
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzć żądania: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwano kod błędu %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}

func TestSearchByPublicationsCount(t *testing.T) {
	router := TestSetUP()

	publicationMinCount := 5
	publicationMaxCount := 20

	url := fmt.Sprintf("http://localhost:8000/api/search?publications_min=%v&publications_max=%v", publicationMinCount, publicationMaxCount)
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result.Scientists {
		bibliometric := GetBibliometrics(item)
		if *bibliometric.PublicationCount < publicationMinCount || *bibliometric.PublicationCount > publicationMaxCount {
			t.Errorf("Niewłaściwa ilość publikacji")
		}
	}
}

func TestSearchByNegativePublicationsCount(t *testing.T) {
	router := TestSetUP()
	publicationMinCount := -1
	publicationMaxCount := -10

	url := fmt.Sprintf("http://localhost:8000/api/search?publications_min=%v&publications_max=%v", publicationMinCount, publicationMaxCount)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwano kodu błędu %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}

func TestSearchByPublicationsCountWhenNotANumber(t *testing.T) {
	router := TestSetUP()
	publicationMinCount := "Pomidor"
	publicationMaxCount := 20

	url := fmt.Sprintf("http://localhost:8000/api/search?publications_min=%v&publications_max=%v", publicationMinCount, publicationMaxCount)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Oczekiwano kodu błędu %d, a otrzymano %d", http.StatusUnprocessableEntity, w.Code)
	}
}

func TestSearchByResearchAreas(t *testing.T) {
	router := TestSetUP()
	researchArea := "information and communication technology (ICT)"

	url := "http://localhost:8000/api/search?research_areas%5B%5D=" + researchArea
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result.Scientists {
		if !ContainsResearchArea(researchArea, item.ResearchAreas) {
			t.Errorf("Naukowiec nie posiada odpowiedniej dyscypliny.")
			t.SkipNow()
		}
	}
}
func TestSearchBySurname(t *testing.T) {
	router := TestSetUP()
	surname := "Kowa"

	url := fmt.Sprintf("http://localhost:8000/api/search?surname=%s", surname)
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}
	for _, item := range result.Scientists {
		if !strings.Contains(strings.ToLower(*item.LastName), strings.ToLower(surname)) {
			t.Errorf("Nazwisko naukowca nie zawiera fragmentu: %s.", surname)
			t.SkipNow()
		}
	}
	t.Logf("Test przeszedł pomyślnie")
}

func TestSearchByBadSurname(t *testing.T) {
	router := TestSetUP()
	surname := "BAD SURNAME"
	url := fmt.Sprintf("http://localhost:8000/api/search?surname=%s", surname)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwno kodu błędu %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}

func TestSearchByYearScoreFilters(t *testing.T) {
	router := TestSetUP()
	year := 2021
	t.Skip()
	url := "http://localhost:8000/api/search?year_score_filters%5B%5D=" + string(rune(year))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %d", err)
	}

	req.Header.Add("Accept", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano kod błędu: %d", w.Code)
	}

	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Otrzymano błda podczas odczytywania odpowiedzi: %d", err)
	}

	var result responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Otrzymano błąda podczas parsowania JSON: %v", err)
	}
	for _, item := range result.Scientists {

		if ContainsPickedYearOfPublication(item.PublicationScores, &year) != nil {
			t.Errorf("Otrzymano kod błędu %s", err.Error())
		}
	}
}
func TestSearchByNegativeYearScoreFilters(t *testing.T) {
	router := TestSetUP()
	year := -2021

	url := "http://localhost:8000/api/search?year_score_filters%5B%5D=" + string(rune(year))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %d", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekwiano kodu błędu %d, a otrzymano %d", http.StatusBadRequest, w.Code)
	}
}
func TestSearchByYearScoreFiltersWhenNotANumber(t *testing.T) {
	router := TestSetUP()
	year := "Pomidor"
	url := "http://localhost:8000/api/search?year_score_filters%5B%5D=" + year
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %d", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Oczekwiano kodu błędu %d, a otrzymano %d", http.StatusUnprocessableEntity, w.Code)
	}
}
