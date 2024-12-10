package tests

import (
	"encoding/json"
	"fmt"
	"io-project-api/internal/models"
	_ "io-project-api/internal/responses"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestRegisterCitationCount(t *testing.T) {

	id := "8611c0f6-039e-4a73-be41-b36ddf4e4674"
	url := fmt.Sprintf("http://localhost:8000/api/bibliometrics/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Otrzymano błąd: %s", res.Status)
	}

	// Wczytaj odpowiedź
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result models.Bibliometrics

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	fmt.Printf("Citation count dla ID %s: %d\n", id, result.CitationCount)
}
