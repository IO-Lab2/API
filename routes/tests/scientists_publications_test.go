package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io-project-api/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterScientistsPublicationsByScientistsID(t *testing.T) {

	router := TestSetUP()

	surname := "Nafkha"
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
	var subject []models.Scientist

	if err := json.Unmarshal(body, &subject); err != nil {
		t.Errorf("Błąd podczas parsowania subject JSON: %v", err)
	}

	if len(subject) == 0 {
		t.Errorf("Nie znaleziono naukowca dla nazwiska %s", surname)
	}

	id := subject[0].ID
	url = fmt.Sprintf("http://localhost:8000/api/scientists_publications/%s", id)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("Nie udało się utworzyć żądania: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if w.Code != http.StatusOK {
		t.Errorf("Otrzymano błąd: %v", w.Code)
	}

	// Wczytaj odpowiedź
	body, err = io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.ScientistPublication

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania result JSON: %v", err)
	}

	t.Logf("Test zakończony pomyślnie.")
}
