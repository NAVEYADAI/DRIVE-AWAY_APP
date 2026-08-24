package sorting_test

import (
	"testing"
	"time"

	"back/internal/sorting"

	"github.com/stretchr/testify/assert"
)

func TestNewChange(t *testing.T) {
	driver := sorting.Driver{
		Id:           5,
		Gps:          "driver-gps",
		TimeOfDriver: time.Date(2026, 1, 1, 6, 0, 0, 0, time.UTC),
	}
	transportation := sorting.Transportation{
		Id:                       2,
		TimeStart:                time.Date(2026, 1, 1, 7, 0, 0, 0, time.UTC),
		GpsEnd:                   "transport-end-gps",
		TimeForAllTransportation: 20 * time.Minute,
	}

	change := sorting.NewChange(driver, transportation)

	assert.Equal(t, 5, change.DriverId)
	assert.Equal(t, 2, change.TransportationId)
	assert.Equal(t, "driver-gps", change.GpsStart, "GpsStart should capture the driver's position before the change")
	assert.Equal(t, "transport-end-gps", change.GpsEnd)
	assert.Equal(t, driver.TimeOfDriver, change.TimeReturnForDriver)
	// TimeToSet = trip start + trip duration + 5 minute buffer
	assert.Equal(t, time.Date(2026, 1, 1, 7, 25, 0, 0, time.UTC), change.TimeToSet)
}

func TestChangeToDriver_GetIdDriverAndTransportation(t *testing.T) {
	tmp := sorting.ChangeToDriver{DriverId: 5, TransportationId: 2}
	driverId, transportationId := tmp.GetIdDriverAndTransportation()
	assert.Equal(t, 5, driverId)
	assert.Equal(t, 2, transportationId)
}

func TestGetTimeTimeInTimeDuration(t *testing.T) {
	cases := []struct {
		name     string
		input    time.Time
		expected time.Duration
	}{
		{"zero", time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), 0},
		{"hours minutes seconds", time.Date(2026, 1, 1, 3, 15, 45, 0, time.UTC), 3*time.Hour + 15*time.Minute + 45*time.Second},
		{"just before midnight", time.Date(2026, 1, 1, 23, 59, 59, 0, time.UTC), 23*time.Hour + 59*time.Minute + 59*time.Second},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, sorting.GetTimeTimeInTimeDuration(c.input))
		})
	}
}
