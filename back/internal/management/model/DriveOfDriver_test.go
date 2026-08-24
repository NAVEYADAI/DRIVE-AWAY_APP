package model_test

import (
	"testing"

	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestCreateDriveOfDriver_And_GetAllByDriverID(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriveOfDriver(1, 10, 101)
	model.CreateDriveOfDriver(1, 20, 101)
	model.CreateDriveOfDriver(1, 30, 202) // different driver, same sort
	model.CreateDriveOfDriver(2, 40, 101) // same driver, different sort

	got := model.GetAllDrivesByDriverID(101, 1)
	require.Len(t, got, 2)
	for _, d := range got {
		require.Equal(t, 101, d.DriverId)
		require.Equal(t, 1, d.SortId)
	}
}

func TestGetAllDrivesByDriverID_NoMatchReturnsEmpty(t *testing.T) {
	testutil.SetupTestDB(t)

	require.Empty(t, model.GetAllDrivesByDriverID(999, 999))
}
