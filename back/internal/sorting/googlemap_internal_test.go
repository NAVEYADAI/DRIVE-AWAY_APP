package sorting

// White-box tests (package sorting, not sorting_test) — needed here because
// parseDistanceMatrixResponse and mapsHTTPClient are unexported test seams
// that exist specifically so this file can exercise FindWay's parsing and
// network-error paths without making a real call to the Google Maps API.

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type failingRoundTripper struct{}

func (failingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("simulated network failure")
}

func TestFindWay_NetworkErrorReturnsZero(t *testing.T) {
	original := mapsHTTPClient
	defer func() { mapsHTTPClient = original }()
	mapsHTTPClient = &http.Client{Transport: failingRoundTripper{}}

	got := FindWay("origin", "destination")

	assert.Equal(t, time.Duration(0), got)
}

func TestParseDistanceMatrixResponse(t *testing.T) {
	cases := []struct {
		name     string
		body     string
		expected time.Duration
	}{
		{
			name:     "valid single-unit duration",
			body:     `{"rows":[{"elements":[{"duration":{"text":"15 mins"}}]}]}`,
			expected: 15 * time.Minute,
		},
		{
			name:     "valid multi-unit duration",
			body:     `{"rows":[{"elements":[{"duration":{"text":"1 hour 5 mins"}}]}]}`,
			expected: 1*time.Hour + 5*time.Minute,
		},
		{name: "malformed json", body: `not json`, expected: 0},
		{name: "empty rows", body: `{"rows":[]}`, expected: 0},
		{name: "missing rows key", body: `{}`, expected: 0},
		{name: "empty elements", body: `{"rows":[{"elements":[]}]}`, expected: 0},
		{name: "missing duration", body: `{"rows":[{"elements":[{}]}]}`, expected: 0},
		{name: "duration text not a string", body: `{"rows":[{"elements":[{"duration":{"text":5}}]}]}`, expected: 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, parseDistanceMatrixResponse([]byte(c.body)))
		})
	}
}

func TestMyApiKey_ReadsEnvVar(t *testing.T) {
	t.Setenv("MY_GOOGLE_API_KEY", "test-key-123")
	assert.Equal(t, "test-key-123", myApiKey())
}

func TestMyApiKey_EmptyWhenUnset(t *testing.T) {
	t.Setenv("MY_GOOGLE_API_KEY", "")
	assert.Equal(t, "", myApiKey())
}
