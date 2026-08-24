package sorting_test

import (
	"testing"
	"time"

	"back/internal/sorting"

	"github.com/stretchr/testify/assert"
)

func TestCreateMap(t *testing.T) {
	drivers := []sorting.Driver{
		{Id: 1, Gps: "a"},
		{Id: 2, Gps: "b"},
	}

	m := sorting.CreateMap(drivers)

	assert.Len(t, m, 2)
	assert.Equal(t, "a", m[1].Gps)
	assert.Equal(t, "b", m[2].Gps)
}

func TestCreateMap_Empty(t *testing.T) {
	m := sorting.CreateMap(nil)
	assert.Empty(t, m)
}

func TestSortFromTime_OrdersByStartTime(t *testing.T) {
	mk := func(id, hour int) sorting.Transportation {
		return sorting.Transportation{Id: id, TimeStart: time.Date(2026, 1, 1, hour, 0, 0, 0, time.UTC)}
	}
	list := []sorting.Transportation{mk(3, 15), mk(1, 6), mk(4, 20), mk(2, 9)}

	sorted := sorting.SortFromTime(list)

	ids := make([]int, len(sorted))
	for i, tr := range sorted {
		ids[i] = tr.Id
	}
	assert.Equal(t, []int{1, 2, 3, 4}, ids)
}

func TestSortFromTime_EmptyList(t *testing.T) {
	sorted := sorting.SortFromTime(nil)
	assert.Empty(t, sorted)
}

func TestSortFromTime_SingleElement(t *testing.T) {
	list := []sorting.Transportation{{Id: 1}}
	sorted := sorting.SortFromTime(list)
	assert.Equal(t, []int{1}, []int{sorted[0].Id})
}

func TestSortFromTime_AlreadySorted(t *testing.T) {
	mk := func(id, hour int) sorting.Transportation {
		return sorting.Transportation{Id: id, TimeStart: time.Date(2026, 1, 1, hour, 0, 0, 0, time.UTC)}
	}
	list := []sorting.Transportation{mk(1, 6), mk(2, 7), mk(3, 8)}

	sorted := sorting.SortFromTime(list)

	ids := make([]int, len(sorted))
	for i, tr := range sorted {
		ids[i] = tr.Id
	}
	assert.Equal(t, []int{1, 2, 3}, ids)
}

func TestSortFromTime_TiedStartTimes(t *testing.T) {
	same := time.Date(2026, 1, 1, 9, 0, 0, 0, time.UTC)
	list := []sorting.Transportation{
		{Id: 1, TimeStart: same},
		{Id: 2, TimeStart: same},
	}

	sorted := sorting.SortFromTime(list)

	assert.Len(t, sorted, 2)
	assert.ElementsMatch(t, []int{1, 2}, []int{sorted[0].Id, sorted[1].Id})
}
