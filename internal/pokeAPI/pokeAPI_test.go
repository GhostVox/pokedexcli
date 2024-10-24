package pokeapi

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Brent-the-carpenter/pokedexcli/internal/pokecache"
	"github.com/Brent-the-carpenter/pokedexcli/types"
)

// Mock HTTP function for testing
func mockPokeAPI() (*http.Response, error) {
	mockResponse := `{
		"count": 300,
		"next": "http://test/20",
		"previous": null,
		"results": [
			{
				"name": "valhalla",
				"url": "https://valhalla"
			},
			{
				"name": "hogwarts",
				"url": "https://hogwarts"
			}
		]
	}`

	// Create a response recorder for testing
	recorder := httptest.NewRecorder()
	recorder.WriteHeader(http.StatusOK)
	io.WriteString(recorder, mockResponse)

	// Create an rt*http.Response from the recorder
	return recorder.Result(), nil
}

// Test the GetLocations function using the mock API
func TestGetLocations(t *testing.T) {
	// Pass the mockPokeAPI function to GetLocations for testing
	var testString string = "Hello nate"
	// Set up a valid config with a cache
	cache := pokecache.NewCache(10 * time.Minute) // Assuming pokecache.NewCache creates a cache
	config := &types.Config{
		Cache: cache, // Initialize the cache
	}
	locations, err := GetLocations(mockPokeAPI, &testString, config)
	if err != nil {
		t.Fatalf("GetLocations returned an error: %v", err)
	}

	// Validate the results
	if len(locations.Results) != 2 {
		t.Errorf("Expected 2 locations, got %d", len(locations.Results))
	}

	if locations.Results[0].Name != "valhalla" {
		t.Errorf("Expected first location to be valhalla, got %s", locations.Results[0].Name)
	}

	if locations.Results[1].Name != "hogwarts" {
		t.Errorf("Expected second location to be hogwarts, got %s", locations.Results[1].Name)
	}
}
