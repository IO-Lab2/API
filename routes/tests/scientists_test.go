package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io-project-api/internal/responses"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRegisterScientists(t *testing.T) {
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

	// Wykonaj zapytanie GET
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
	var result responses.ScientistBody

	if err := json.Unmarshal(body, &result); err != nil {
		t.Errorf("Błąd podczas parsowania result JSON: %v", err)
	}

	expected := subject.Scientists[0]
	received := result
	if expected.ID != received.ID {
		t.Errorf("Różne ID: oczekiwano %s, otrzymano %s", expected.ID, received.ID)
	}
	if expected.FirstName != received.FirstName {
		t.Errorf("Różne imiona: oczekiwano %s, otzymano %s", expected.FirstName, received.FirstName)
	}
	if expected.LastName != received.LastName {
		t.Errorf("Różne nazwiska: oczekiwano %s, otrzymano %s", expected.LastName, received.LastName)
	}
	if expected.AcademicTitle != received.AcademicTitle {
		t.Errorf("Różne tytuły: oczekiwano %s, otrzymano %s", expected.AcademicTitle, received.AcademicTitle)
	}
	if *expected.Position != *received.Position {
		t.Errorf("Różne pozycje: oczekwiano %v, otrzymano %v", *expected.Position, *received.Position)
	}
	if !reflect.DeepEqual(expected.ResearchAreas, received.ResearchAreas) {
		t.Errorf("Różne dyscypliny: oczekiwano %+v, otrzymano %+v", expected.ResearchAreas, received.ResearchAreas)
	}
	if *expected.Email != *received.Email {
		t.Errorf("Rózne maile: oczekiwano %s, otrzymano %s", *expected.Email, *received.Email)
	}
	if *expected.ProfileUrl != *received.ProfileUrl {
		t.Errorf("Różne hiperłącza: oczekiwano %s, otrzymano %s", *expected.ProfileUrl, *received.ProfileUrl)
	}
	if expected.CreatedAt != received.CreatedAt {
		t.Errorf("Różne czasy utworzenia: oczekiwano %s, otrzymano %s", expected.CreatedAt, received.CreatedAt)
	}
	if expected.UpdatedAt != received.UpdatedAt {
		t.Errorf("Różne czasy uaktualniania: oczekiawno %s, otrzymano %s", expected.UpdatedAt, received.UpdatedAt)
	}
	if !reflect.DeepEqual(expected.Bibliometrics, received.Bibliometrics) {
		t.Errorf("Różne bibliometrie: oczekiwano %+v, otzrmano %+v", expected.Bibliometrics, received.Bibliometrics)
	}
	err = ComparePublicationScores(expected.PublicationScores, received.PublicationScores)
	if err != nil {
		t.Errorf(err.Error())
	}
}
