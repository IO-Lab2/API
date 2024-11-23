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

func TestRegisterMinisterialScore(t *testing.T) {

	id := "cd85ed8e-4c50-45c7-90dd-24d34323ee74"
	url := fmt.Sprintf("http://127.0.0.1:8000/api/bibliometrics/%s", id)

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
	var result []models.Bibliometrics

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	fmt.Printf("Ministerial score dla ID %s: %f\n", id, result[0].MinisterialScore)
}
