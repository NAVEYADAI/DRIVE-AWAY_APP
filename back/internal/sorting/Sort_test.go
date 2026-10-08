package sorting_test

import (
	"testing"

	"back/internal/sorting"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Note: this file previously contained tests for back/internal/management's
// model package (CreateSort, GetLastSortByCompanyId, etc.) reached through
// config.Connect() directly against whatever DB happened to be configured —
// that's both the wrong package to be testing from here and exactly the
// "test hits the wrong/uncontrolled DB" pattern the new testing policy
// forbids. Those functions are back/internal/management's responsibility to
// test, not sorting's. This file now actually tests sorting.Sort().

func TestSort_EmptyTransportationList_ReturnsEmptyScheduleWithoutError(t *testing.T) {
	drivers := []sorting.Driver{sorting.NewDriver(1, "10,10", false)}

	schedule, message := sorting.Sort(drivers, nil)

	require.Equal(t, "good work", message)
	assert.Equal(t, 0, schedule.LenOfTransportation(1))
}

func TestSort_NoDriversNoTransportations_ReturnsEmptySchedule(t *testing.T) {
	schedule, message := sorting.Sort(nil, nil)

	require.Equal(t, "good work", message)
	assert.NotNil(t, message)
	_ = schedule // nil TravelListForDrivers is a valid, usable zero value
}

// TestSort_AssignsTripsToDrivers is a full integration-style test of the
// scheduling algorithm, but it's skipped: Sort -> Grade -> FindWay makes a
// real HTTP call to the Mapbox API for every driver/transportation
// pair, which is unsuitable for a unit test (network-dependent, requires a
// real API key, non-deterministic). See MapboxApi_test.go for coverage of the pieces that can be tested
// without the network. Making Sort's distance lookup injectable (so this
// could run against a fake) would be a reasonable follow-up but is a real
// behavioral/API change, out of scope for a test-only pass.
func TestSort_AssignsTripsToDrivers(t *testing.T) {
	t.Skip("Sort() requires a real Mapbox API call via Grade->FindWay; not unit-testable without a network/API-key dependency or a Sort-level refactor to inject the distance lookup")
}
