package _MainSort_test_test

import (
	"MainSorting/MainSort"
	"testing"
)

var myHome = "31.787824575497734, 34.760588310287254"
var mySchool = "31.792839967395615, 34.818742093189805"
var timeToDrive = 15

func TestTime(t *testing.T) {
	if int(MainSort.FindWay(myHome, mySchool).Minutes()) == 15 {
		println("work")
	} else {
		t.Error("error in time")
	}
}
