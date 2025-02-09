package _MainSort_test_test

import (
	"MainSorting/MainSort"
	"container/heap"
	"testing"
)

func TestLen(t *testing.T) {
	var tmp MainSort.MinHeap
	heap.Init(&tmp)
	driver := []MainSort.Driver{
		{Id: 1, Gps: "10,23"},
		{Id: 2, Gps: "53.56"},
		{Id: 3, Gps: "65,78"},
		{Id: 4, Gps: "90,43"},
	}
	tmp.Push(MainSort.Item{Grate: 50, Driver: driver[0]})
	tmp.Push(MainSort.Item{Grate: 22, Driver: driver[2]})
	tmp.Push(MainSort.Item{Grate: 48, Driver: driver[3]})
	tmp.Push(MainSort.Item{Grate: 70, Driver: driver[1]})

	if heap.Pop(&tmp).(MainSort.Item).Driver.Id != 3 {
		t.Error("error in out form heap ")
	}
	if heap.Pop(&tmp).(MainSort.Item).Driver.Id != 4 {
		t.Error("error in out 2 form heap ")
	}
	if heap.Pop(&tmp).(MainSort.Item).Driver.Id != 1 {
		t.Error("error in out 3 form heap ")
	}
	println("work")
}
