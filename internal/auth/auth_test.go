package auth

import (
	"net/http"
    "testing"
)

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := http.Header{}

	apiKey, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected error '%v', got '%v'", ErrNoAuthHeaderIncluded, err)
	}
	if apiKey != "" {
		t.Errorf("expected empty apiKey, got '%s'", apiKey)
	}
}
