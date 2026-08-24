package sorting_test

import (
	"testing"

	"back/internal/sorting"

	"github.com/stretchr/testify/assert"
)

func TestTravelListForDrivers_AddDrive_InitializesNilMap(t *testing.T) {
	var tmp sorting.TravelListForDrivers
	tmp.AddDrive(1, 1)
	tmp.AddDrive(2, 1)
	tmp.AddDrive(3, 2)

	assert.Equal(t, 2, tmp.LenOfTransportation(1))
	assert.Equal(t, 1, tmp.LenOfTransportation(2))
}

func TestTravelListForDrivers_AddDrive_MultipleDrivers(t *testing.T) {
	var tmp sorting.TravelListForDrivers
	tmp.AddDrive(1, 1)
	tmp.AddDrive(2, 1)
	tmp.AddDrive(3, 1)
	tmp.AddDrive(4, 2)
	tmp.AddDrive(5, 2)

	assert.Equal(t, 3, tmp.LenOfTransportation(1))
	assert.Equal(t, 2, tmp.LenOfTransportation(2))
}

func TestTravelListForDrivers_RemoveLastTransportation(t *testing.T) {
	var tmp sorting.TravelListForDrivers
	tmp.AddDrive(1, 1)
	tmp.AddDrive(2, 1)
	tmp.AddDrive(3, 1)

	tmp.RemoveLastTransportation(1)
	assert.Equal(t, 2, tmp.LenOfTransportation(1))

	tmp.RemoveLastTransportation(1)
	assert.Equal(t, 1, tmp.LenOfTransportation(1))
}

func TestTravelListForDrivers_RemoveLastTransportation_OnEmptyDriver_NoPanic(t *testing.T) {
	var tmp sorting.TravelListForDrivers
	assert.NotPanics(t, func() {
		tmp.RemoveLastTransportation(999)
	})
	assert.Equal(t, 0, tmp.LenOfTransportation(999))
}

func TestTravelListForDrivers_LenOfTransportation_UnknownDriver(t *testing.T) {
	var tmp sorting.TravelListForDrivers
	tmp.AddDrive(1, 1)
	assert.Equal(t, 0, tmp.LenOfTransportation(2), "a driver with no assigned trips should report 0, not panic or a stale count")
}

func TestTravelListForDrivers_NilMap_LenOfTransportationIsZero(t *testing.T) {
	var tmp sorting.TravelListForDrivers
	assert.Equal(t, 0, tmp.LenOfTransportation(1))
}
