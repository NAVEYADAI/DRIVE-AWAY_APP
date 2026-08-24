package sorting_test

import (
	"testing"
	"time"

	"back/internal/sorting"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDriver(t *testing.T) {
	d := sorting.NewDriver(7, "10,20", true)
	assert.Equal(t, 7, d.Id)
	assert.Equal(t, "10,20", d.Gps)
	assert.True(t, d.WageEarner)
	require.Equal(t, 6, d.TimeOfDriver.Hour(), "a fresh driver's clock should start at 06:00")
	assert.Equal(t, 0, d.TimeOfDriver.Minute())
}

func TestDriver_GetId(t *testing.T) {
	d := sorting.Driver{Id: 42}
	assert.Equal(t, 42, d.GetId())
}

func TestDriver_AddChangeFromDriver_MovesDriverToDestination(t *testing.T) {
	d := sorting.Driver{Id: 1, Gps: "start-gps", TimeOfDriver: time.Date(2026, 1, 1, 6, 0, 0, 0, time.UTC)}
	change := sorting.ChangeToDriver{
		GpsStart:            "start-gps",
		GpsEnd:              "end-gps",
		TimeToSet:           time.Date(2026, 1, 1, 8, 0, 0, 0, time.UTC),
		TimeReturnForDriver: time.Date(2026, 1, 1, 6, 0, 0, 0, time.UTC),
	}

	d.AddChangeFromDriver(change)

	assert.Equal(t, "end-gps", d.Gps, "applying a change should move the driver to where the trip ended")
	assert.Equal(t, change.TimeToSet, d.TimeOfDriver)
}

// TestDriver_MoveChangeFromDriver_RevertsGpsToOriginalPosition is a regression
// test for a bug where undoing a driver's assignment (used by the Sort
// algorithm's backtracking) left the driver's GPS at the trip's destination
// instead of restoring it to where the driver was before the trip was
// assigned. MoveChangeFromDriver wrote change.GpsEnd (the *new* position,
// identical to what AddChangeFromDriver writes) instead of change.GpsStart
// (the position saved before the change was applied), so after a backtrack
// the driver's location was silently wrong for any subsequent Grade()
// distance calculation.
func TestDriver_MoveChangeFromDriver_RevertsGpsToOriginalPosition(t *testing.T) {
	d := sorting.Driver{Id: 1, Gps: "end-gps", TimeOfDriver: time.Date(2026, 1, 1, 8, 0, 0, 0, time.UTC)}
	change := sorting.ChangeToDriver{
		GpsStart:            "start-gps",
		GpsEnd:              "end-gps",
		TimeToSet:           time.Date(2026, 1, 1, 8, 0, 0, 0, time.UTC),
		TimeReturnForDriver: time.Date(2026, 1, 1, 6, 0, 0, 0, time.UTC),
	}

	d.MoveChangeFromDriver(change)

	assert.Equal(t, "start-gps", d.Gps, "reverting a change must restore the driver's original GPS position")
	assert.Equal(t, change.TimeReturnForDriver, d.TimeOfDriver)
}
