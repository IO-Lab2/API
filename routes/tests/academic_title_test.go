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

	"github.com/google/uuid"
)

func TestRegisterAcademicTitle(t *testing.T) {

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
	url = fmt.Sprintf("http://localhost:8000/api/scientists/%s", id)

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
	var result models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania JSON: %v", err)
	}

	// Porównanie z tytułów
	if *subject.Scientists[0].AcademicTitle != result.AcademicTitle {
		t.Errorf("Tytułu naukowców się rożnią. Oczekiwano: %+v, Otrzymano: %+v", subject.Scientists[0].AcademicTitle, result.AcademicTitle)
	}
}
func TestRegisterAcademicTitleBadID(t *testing.T) {
	router := TestSetUP()
	id := uuid.New()
	url := fmt.Sprintf("http://localhost:8000/api/scientists/%s", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nieudało się utworzyć żądania %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Oczekiwano kod błedu 400. Otrzymano %d", w.Code)
	}
}
func TestRegisterAcademicTitleIDNil(t *testing.T) {
	router := TestSetUP()
	id := ""
	url := fmt.Sprintf("http://localhost:8000/api/scientists/%s", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Nieudało się utworzyć żądania %v", err)
	}
	req.Header.Add("Accept", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Errorf("Oczekiwano kod błedu %d. Otrzymano %d", http.StatusBadRequest, w.Code)
	}
}
