package _MainSort_test_test

import (
	"MainSorting/MainSort"
	"log"
	"testing"
)

func TestAddAnd3Travel(t *testing.T) {
	var tmp MainSort.TravelListForDrivers
	tmp.AddDrive(1, 1)
	tmp.AddDrive(2, 1)
	tmp.AddDrive(3, 2)
	if tmp.LenOfTransportation(1) != 3 {
		log.Println("error in loading drive for Driver")
	} else {
		println("work")
	}
}
func TestAddFor2Driver(t *testing.T) {
	var tmp MainSort.TravelListForDrivers

	tmp.AddDrive(1, 1)
	tmp.AddDrive(2, 1)
	tmp.AddDrive(3, 1)

	tmp.AddDrive(4, 2)
	tmp.AddDrive(5, 2)
	if tmp.LenOfTransportation(1) != 3 && tmp.LenOfTransportation(2) != 2 {
		log.Println("error in loading drive for Driver")
	} else {
		println("work")
	}
}

func TestAddAndRemove(t *testing.T) {
	var tmp MainSort.TravelListForDrivers

	tmp.AddDrive(1, 1)
	tmp.AddDrive(2, 1)
	tmp.AddDrive(3, 1)
	tmp.RemoveLastTransportation(1)
	tmp.RemoveLastTransportation(1)
	tmp.AddDrive(4, 2)
	tmp.AddDrive(5, 2)
	if tmp.LenOfTransportation(1) != 1 && tmp.LenOfTransportation(2) != 2 {
		log.Println("error in loading drive for Driver")
	} else {
		println("work")
	}
}

func TestAddAndRemoveNotIn(t *testing.T) {
	var tmp MainSort.TravelListForDrivers

	tmp.AddDrive(1, 1)
	tmp.AddDrive(2, 1)
	tmp.AddDrive(3, 1)
	tmp.RemoveLastTransportation(1)
	tmp.AddDrive(4, 2)
	tmp.RemoveLastTransportation(1)
	tmp.AddDrive(5, 2)
	if tmp.LenOfTransportation(1) != 1 && tmp.LenOfTransportation(2) != 2 {
		log.Println("error in loading drive for Driver")
	} else {
		println("work")
	}
}
