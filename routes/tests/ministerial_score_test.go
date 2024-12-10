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
		log.Fatalf("Revived error: %s", res.Status)
	}

	// Wczytaj odpowiedź
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error occured while reading response: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result models.Bibliometrics

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error occured during deserialization process: %v", err)
	}

	fmt.Printf("Ministerial score for ID %s: %f\n", id, result.MinisterialScore)
}
