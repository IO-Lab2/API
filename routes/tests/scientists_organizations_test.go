package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io-project-api/internal/models"
	"io-project-api/internal/responses"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterScientistsOrganizationsByID(t *testing.T) {

	router := TestSetUP()

	id := "d58b4cf2-f79b-4820-a465-868892e122a6"
	url := fmt.Sprintf("http://localhost:8000/api/scientists_organizations/%s", id)

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
	var result []models.ScientistOrganization

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania result JSON: %v", err)
	}

	t.Logf("Test zakończony pomyślnie.")
}
func TestRegisterScientistsOrganizationsByScientistID(t *testing.T) {

	router := TestSetUP()
	name := "Marcin"
	surname := "Bator"
	url := fmt.Sprintf("http://localhost:8000/api/search?name=%s&surname=%s", name, surname)

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
	var subject *responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &subject); err != nil {
		t.Errorf("Błąd podczas parsowania subject JSON: %v", err)
	}

	if subject == nil {
		t.Errorf("Nie znaleziono naukowca dla imienia: %s i nazwiska: %s", name, surname)
	}

	id := subject.Scientists[0].ID
	url = fmt.Sprintf("http://localhost:8000/api/organizations/scientist/%s", id)

	req, err = http.NewRequest(http.MethodGet, url, nil)
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
	var result []models.ScientistOrganization

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania result JSON: %v", err)
	}

	t.Logf("Test zakończony pomyślnie.")
}
