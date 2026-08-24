package sorting_test

import (
	"testing"
	"time"

	"back/internal/sorting"

	"github.com/stretchr/testify/assert"
)

// FindWay itself (the function that actually calls the Google Maps API) is
// intentionally not exercised here — a real unit test must not depend on
// network access or a real API key. Its network-error path and its response
// parsing (parseDistanceMatrixResponse) are covered as white-box tests in
// googlemap_internal_test.go instead. This file only covers the exported,
// pure GetInTime duration parser.
func TestGetInTime(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected time.Duration
	}{
		{"minutes only", "15 mins", 15 * time.Minute},
		{"single minute singular unit", "1 min", time.Minute},
		{"hours and minutes", "1 hour 5 mins", 1*time.Hour + 5*time.Minute},
		{"days hours minutes", "1 day 2 hours 3 mins", 24*time.Hour + 2*time.Hour + 3*time.Minute},
		{"days plural", "2 days", 48 * time.Hour},
		{"empty string", "", 0},
		{"unrecognized unit", "10 fortnights", 0},
		// Regression test: GetInTime used to panic with "index out of range"
		// on an odd number of whitespace tokens (a trailing value with no
		// unit) instead of returning 0 like every other malformed-input
		// path in this function. A crafted or unexpected Google Maps API
		// response text would have crashed the request-handling goroutine.
		{"malformed odd token count (regression, used to panic)", "10", 0},
		{"malformed odd token count, multiple pairs", "1 hour 5", 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, sorting.GetInTime(c.input))
		})
	}
}
