package tests

import (
	"encoding/json"
	"fmt"
	"io-project-api/internal/models"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestRegisterResearchTitle(t *testing.T) {

	id := "72e6d858-222b-48c0-819a-4c81081c787b"
	url := fmt.Sprintf("http://127.0.0.1:8000/api/scientists/%s", id)

	// Wykonaj zapytanie GET
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Błąd podczas wysyłania zapytania: %v", err)
	}
	defer resp.Body.Close()

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Otrzymano błąd: %s", resp.Status)
	}

	// Wczytaj odpowiedź
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	fmt.Printf("Research title dla ID %s: %s\n", id, result[0].ResearchArea)
}
