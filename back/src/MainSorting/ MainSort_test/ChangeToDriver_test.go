package _MainSort_test_test

import (
	"MainSorting/MainSort"
	"testing"
	"time"
)

func TestGetIDDriverAndDrive(t *testing.T) {
	tmp := MainSort.ChangeToDriver{
		DriverId:         5,
		TransportationId: 2,
		TimeToSet:        time.Hour,
		GpsStart:         "10,30",
		GpsEnd:           "30,10",
	}
	x, y := tmp.GetIdDriverAndTransportation()
	if x != 5 || y != 2 {
		t.Error("this func didnt work")
	}
	print("work")
}
