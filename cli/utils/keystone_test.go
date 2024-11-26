package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendSlackNotification_Success(t *testing.T) {
	// Set up environment variables
	t.Setenv("GITHUB_SERVER_URL", "https://github.com")
	t.Setenv("GITHUB_REPOSITORY", "owner/repo")
	t.Setenv("GITHUB_RUN_ID", "123456")
	t.Setenv("GITHUB_ACTOR", "test-actor")

	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method and content type
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Decode the request body
		var payload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			t.Errorf("Error decoding JSON payload: %v", err)
		}

		// Verify the payload structure
		blocks, ok := payload["blocks"].([]interface{})
		if !ok || len(blocks) != 2 {
			t.Errorf("Expected 'blocks' with 2 sections, got %v", payload["blocks"])
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	// Call the function
	err := SendSlackNotification(mockServer.URL, "test-command")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSendSlackNotification_NonOKResponse(t *testing.T) {
	// Set up environment variables
	t.Setenv("GITHUB_SERVER_URL", "https://github.com")
	t.Setenv("GITHUB_REPOSITORY", "owner/repo")
	t.Setenv("GITHUB_RUN_ID", "123456")
	t.Setenv("GITHUB_ACTOR", "test-actor")

	// Create a mock server that returns 500
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	// Call the function
	err := SendSlackNotification(mockServer.URL, "test-command")
	if err == nil {
		t.Error("Expected error due to non-OK response, got nil")
	}
}

func TestSendSlackNotification_HTTPError(t *testing.T) {
	// Set up environment variables
	t.Setenv("GITHUB_SERVER_URL", "https://github.com")
	t.Setenv("GITHUB_REPOSITORY", "owner/repo")
	t.Setenv("GITHUB_RUN_ID", "123456")
	t.Setenv("GITHUB_ACTOR", "test-actor")

	// Use an invalid URL to simulate HTTP error
	invalidURL := "http://invalid.url"

	// Call the function
	err := SendSlackNotification(invalidURL, "test-command")
	if err == nil {
		t.Error("Expected error due to HTTP failure, got nil")
	}
}
