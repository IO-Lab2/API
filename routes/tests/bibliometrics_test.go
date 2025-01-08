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

func TestRegisterBibliometricsRoutesByID(t *testing.T) {

	router := TestSetUP()

	// Przypisujemy zmienną ID
	id := "8611c0f6-039e-4a73-be41-b36ddf4e4674"
	url := fmt.Sprintf("http://localhost:8000/api/bibliometrics/%s", id)

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
	var result models.Bibliometrics

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}

	t.Logf("Test zakończony pomyślnie.")

}
func TestRegisterBibliometricsRoutesByAuthor(t *testing.T) {

	router := TestSetUP()
	name := "Marcin"
	surname := "Bator"
	url := fmt.Sprintf("http://localhost:8000/api/search?name=%s&surname=%s", name, surname)

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
	var subject *responses.ScientistsResponseBody

	if err := json.Unmarshal(body, &subject); err != nil {
		t.Errorf("Błąd podczas parsowania subject JSON: %v", err)
	}

	if subject == nil {
		t.Errorf("Nie znaleziono naukowca dla imienia: %s i nazwiska: %s", name, surname)
	}

	id := subject.Scientists[0].ID
	url = fmt.Sprintf("http://localhost:8000/api/bibliometrics/author/%s", id)

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
	var result models.Bibliometrics

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania result JSON: %v", err)
	}

	t.Logf("Test zakończony pomyślnie.")
}
