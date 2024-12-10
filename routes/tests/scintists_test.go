package tests

import (
	"encoding/json"
	"fmt"
	"io-project-api/internal/models"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestRegisterScientists(t *testing.T) {

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
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	if len(subject) == 0 {
		t.Fatalf("Nie znaleziono naukowca dla nazwiska %s", surname)
	}

	id := subject[0].ID
	url = fmt.Sprintf("http://localhost:8000/api/scientists/%s", id)

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
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Błąd podczas odczytywania odpowiedzi: %v", err)
	}

	// Rozpakuj JSON do struktury
	var result []models.Scientist

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Błąd podczas parsowania JSON: %v", err)
	}

	// Porównanie z `reflect.DeepEqual`
	if !reflect.DeepEqual(subject[0], result[0]) {
		t.Fatalf("Dane naukowców różnią się. Oczekiwano: %+v, Otrzymano: %+v", subject, result)
	}

	t.Logf("Test zakończony pomyślnie, dane naukowców są zgodne.")
}
