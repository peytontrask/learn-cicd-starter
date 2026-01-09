package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case: No Authorization header
	_, err := GetAPIKey(http.Header{})
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}

	_, err = GetAPIKey(http.Header{"Authorization": []string{"InvalidHeader"}})
	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("expected malformed header error, got %v", err)
	}

	// Test case: Valid Authorization header
	apiKey, err := GetAPIKey(http.Header{"Authorization": []string{"ApiKey my-secret-key"}})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if apiKey != "my-secret-key" {
		t.Errorf("expected api key 'my-secret-key', got '%s'", apiKey)
	}
}
