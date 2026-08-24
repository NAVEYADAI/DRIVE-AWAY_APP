package sorting_test

import (
	"container/heap"
	"testing"

	"back/internal/sorting"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMinHeap_PopsInAscendingGradeOrder(t *testing.T) {
	var tmp sorting.MinHeap
	heap.Init(&tmp)
	driver := []sorting.Driver{
		{Id: 1, Gps: "10,23"},
		{Id: 2, Gps: "53,56"},
		{Id: 3, Gps: "65,78"},
		{Id: 4, Gps: "90,43"},
	}
	heap.Push(&tmp, sorting.Item{Grate: 50, Driver: driver[0]})
	heap.Push(&tmp, sorting.Item{Grate: 22, Driver: driver[2]})
	heap.Push(&tmp, sorting.Item{Grate: 48, Driver: driver[3]})
	heap.Push(&tmp, sorting.Item{Grate: 70, Driver: driver[1]})

	require.Equal(t, 4, tmp.Len())

	order := []int{}
	for tmp.Len() > 0 {
		order = append(order, heap.Pop(&tmp).(sorting.Item).Driver.Id)
	}

	assert.Equal(t, []int{3, 4, 1, 2}, order, "items must come out in ascending Grate order (3:22, 4:48, 1:50, 2:70)")
}

func TestMinHeap_Len_EmptyHeap(t *testing.T) {
	var tmp sorting.MinHeap
	assert.Equal(t, 0, tmp.Len())
}

func TestMinHeap_SingleItem(t *testing.T) {
	var tmp sorting.MinHeap
	heap.Init(&tmp)
	heap.Push(&tmp, sorting.Item{Grate: 1, Driver: sorting.Driver{Id: 9}})

	require.Equal(t, 1, tmp.Len())
	item := heap.Pop(&tmp).(sorting.Item)
	assert.Equal(t, 9, item.Driver.Id)
	assert.Equal(t, 0, tmp.Len())
}

func TestMinHeap_TiedGrades_BothComeOut(t *testing.T) {
	var tmp sorting.MinHeap
	heap.Init(&tmp)
	heap.Push(&tmp, sorting.Item{Grate: 5, Driver: sorting.Driver{Id: 1}})
	heap.Push(&tmp, sorting.Item{Grate: 5, Driver: sorting.Driver{Id: 2}})

	seen := map[int]bool{}
	seen[heap.Pop(&tmp).(sorting.Item).Driver.Id] = true
	seen[heap.Pop(&tmp).(sorting.Item).Driver.Id] = true
	assert.True(t, seen[1])
	assert.True(t, seen[2])
}
