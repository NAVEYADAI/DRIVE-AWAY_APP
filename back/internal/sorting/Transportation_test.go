package sorting_test

import (
	"testing"
	"time"

	"back/internal/sorting"

	"github.com/stretchr/testify/assert"
)

func TestGetTimeFromHourAndMin(t *testing.T) {
	cases := []struct {
		hour, minute int
		expected     time.Duration
	}{
		{0, 0, 0},
		{1, 30, 90 * time.Minute},
		{0, 10, 10 * time.Minute},
		{5, 0, 5 * time.Hour},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, sorting.GetTimeFromHourAndMin(c.hour, c.minute))
	}
}

func TestTransportation_GetFinishTransportation(t *testing.T) {
	// Built directly rather than via NewTransportation, which calls the real
	// Google Maps API through FindWay — see GoogleMapApi_test.go for why
	// that call isn't exercised in unit tests.
	transportation := sorting.Transportation{
		TimeStart:                time.Date(2026, 1, 1, 7, 0, 0, 0, time.UTC),
		TimeForAllTransportation: 20 * time.Minute,
	}

	finish := transportation.GetFinishTransportation()

	// start + TimeToDisperseThem (10 min) + trip duration (20 min)
	assert.Equal(t, time.Date(2026, 1, 1, 7, 30, 0, 0, time.UTC), finish)
}
