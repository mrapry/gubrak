package gubrak

import (
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := LoadConfig("example.config.json")

	if err != nil {
		t.Error("load config should not error")
	}

	expectedBaseURL := "http://example.com"
	expectedAuthorization := "Basic exnfekeoeoeojsjalaljahhd"

	if config.URL != expectedBaseURL {
		t.Error("expectedBaseURL should equal config.BaseURL")
	}

	if config.Headers["Authorization"] != expectedAuthorization {
		t.Error("expectedAuthorization should equal headers.Authorization")
	}
}
