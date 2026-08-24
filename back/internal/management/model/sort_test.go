package model_test

import (
	"testing"

	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestCreateSort_And_GetLastSortByCompanyId(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateSort(1, "Sort A")
	model.CreateSort(1, "Sort B")
	model.CreateSort(2, "Other Company Sort")

	last := model.GetLastSortByCompanyId(1)
	require.Equal(t, "Sort B", last.Name, "should return the most recently created sort for that company")
}

func TestGetNumOfLastSortByCompanyId_UnknownCompanyReturnsZero(t *testing.T) {
	testutil.SetupTestDB(t)

	require.Zero(t, model.GetNumOfLastSortByCompanyId(999))
}

func TestGetNumOfLastSortByCompanyId_ReturnsRowId(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateSort(1, "Sort A")
	last := model.GetLastSortByCompanyId(1)

	require.Equal(t, last.ID, model.GetNumOfLastSortByCompanyId(1))
}
