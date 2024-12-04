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

func TestRegisterScientistsOrganizations(t *testing.T) {

	id := "4303b618-6dad-45c4-8917-86c222e5223d"
	url := fmt.Sprintf("http://127.0.0.1:8000/api/scientists_organizations/%s", id)

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
	var result models.ScientistOrganization

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	fmt.Printf("Organization dla ID %s: %s\n", result.ID, result)
}
