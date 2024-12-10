package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io-project-api/internal/models"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestRegisterScientistsOrganizationsByID(t *testing.T) {

	id := "d58b4cf2-f79b-4820-a465-868892e122a6"
	url := fmt.Sprintf("http://localhost:8000/api/scientists_organizations/%s", id)

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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

	// Rozpakuj JSON do struktury
	var result []models.ScientistOrganization

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	fmt.Printf("Organization dla ID %s: %s\n", result[0].ID, result[0])
}
func TestRegisterScientistsOrganizationsByScientistID(t *testing.T) {
	surname := "Bator"
	url := fmt.Sprintf("http://localhost:8000/api/search?surname=%s", surname)

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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var subject []models.Scientist

	if err := json.Unmarshal(body, &subject); err != nil {
		log.Fatalf("(Subject) Błąd podczas parsowania JSON: %v", err)
	}

	if len(subject) == 0 {
		t.Fatalf("Nie znaleziono naukowca dla nazwiska %s", surname)
	}

	id := subject[0].ID
	url = fmt.Sprintf("http://localhost:8000/api/organizations/scientist/%s", id)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

	// Rozpakuj JSON do struktury
	var result []models.ScientistOrganization

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}
}
