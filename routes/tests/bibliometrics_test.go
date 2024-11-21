package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io-project-api/internal/models"
	"log"
	"net/http"
	"testing"
)

func TestRegisterBibliometricsRoutes(t *testing.T) {

	// Przypisujemy zmienną ID
	id := "cd85ed8e-4c50-45c7-90dd-24d34323ee74"
	url := fmt.Sprintf("http://127.0.0.1:8000/api/bibliometrics/%s", id)

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
	var result []models.Bibliometrics

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	fmt.Printf("Bibliometrics dla ID %s:", id)

	for _, bibliometric := range result {
		fmt.Println("\n", bibliometric)
	}

}
