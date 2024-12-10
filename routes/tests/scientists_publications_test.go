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

func TestRegisterScientistsPublicationsByID(t *testing.T) {

	id := "e3efe632-fd8a-4180-a2a3-bc99037fa45a"
	url := fmt.Sprintf("http://127.0.0.1:8000/api/scientists_publications/%s", id)

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
	var result []models.ScientistPublication

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	fmt.Printf("Scientists publications dla ID: %s  %s\n", result[0].ID, result[0])
}
func TestRegisterScientistsPublicationsByScientistsID(t *testing.T) {
	surname := "Nafkha"
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
	url = fmt.Sprintf("http://localhost:8000/api/scientists_publications/%s", id)

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
		log.Fatalf("Otrzymano błąd: %s", res.Status)
	}

	// Wczytaj odpowiedź
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.ScientistPublication

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	fmt.Printf("Scientists publications dla ID: %s  %s\n", result[0].ID, result[0])
}
