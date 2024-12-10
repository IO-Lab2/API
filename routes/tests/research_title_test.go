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
		log.Fatalf("Error occured while getting resposne: %v", err)
	}

	// Rozpakuj JSON do struktury
	var subject []models.Scientist

	if err := json.Unmarshal(body, &subject); err != nil {
		log.Fatalf("Error occured during deserialization process: %v", err)
	}

	if len(subject) == 0 {
		t.Fatalf("Scientist with this surname not found %s", surname)
	}

	id := subject[0].ID
	url = fmt.Sprintf("http://localhost:8000/api/scientists/%s", id)

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

	// Sprawdź, czy zapytanie zakończyło się sukcesem
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Got error: %s", res.Status)
	}

	// Wczytaj odpowiedź
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error occured while getting resposne: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error occured during deserialization process: %v", err)
	}

	if subject[0].ResearchArea != result[0].ResearchArea {
		t.Fatalf("Research area differs. Expected: %+v, Got: %+v", subject[0].AcademicTitle, result[0].AcademicTitle)
	}

	t.Logf("Test finished successfully.")
}
