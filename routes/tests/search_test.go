package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io-project-api/internal/database"
	"io-project-api/internal/models"
	"io-project-api/internal/services"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchAcademicTitle(t *testing.T) {

	router := TestSetUP()

	url := "http://localhost:8000/api/search?academic_titles%5B%5D=PhD"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
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
			t.Errorf("AcademicTitle is wrong.")
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
		t.Errorf("Failed to create request: %v", err)
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
		database.InitDB()
		if !GetBibliometrics(item, minimalScore, maximalScore) {
			t.Errorf("Bibliometrics is wrong.")
		}
		database.CloseDB()
	}
}
func GetBibliometrics(scientists models.Scientist, min float64, max float64) bool {

	TestSetUP()

	result, err := services.GetBibliometricByScientistID(scientists.ID)
	if err != nil {
		log.Fatalf("GetBibliometrics error: %v", err)

		return false
	}
	if result == nil {
		log.Fatalf("GetBibliometrics result is nil")
		return false
	}
	if result.MinisterialScore < min || result.MinisterialScore > max {
		log.Fatalf("MinisterialScore is wrong.")
		return false
	}
	return true
}
func TestSearchByName(t *testing.T) {

	router := TestSetUP()

	name := "Marcin"
	url := fmt.Sprintf("http://localhost:8000/api/search?name=%s", name)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Failed to create request: %v", err)
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
			t.Errorf("FirstName is wrong.")
		}
	}
}
