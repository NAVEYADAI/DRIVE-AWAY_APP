package model_test

import (
	"testing"

	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestCreateDriver_And_GetDriverById(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriver(model.NewDriver(101, 1, "1 Driver St", true))

	got := model.GetDriverById(101)
	require.Equal(t, 1, got.CompanyId)
	require.True(t, got.IsEmployee)
}

func TestGetDriverById_LooksUpByUserId(t *testing.T) {
	testutil.SetupTestDB(t)

	// GetDriverById filters on "userid", not the row's own primary key —
	// confirm it resolves by user id, not by row id.
	model.CreateDriver(model.NewDriver(101, 1, "addr-a", true))
	model.CreateDriver(model.NewDriver(102, 1, "addr-b", false))

	got := model.GetDriverById(102)
	require.Equal(t, "addr-b", got.Address)
	require.False(t, got.IsEmployee)
}

func TestGetAllDriversInCompany_FiltersByCompany(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriver(model.NewDriver(101, 1, "addr-a", true))
	model.CreateDriver(model.NewDriver(102, 2, "addr-b", true))

	got := model.GetAllDriversInCompany(1)
	require.Len(t, got, 1)
	require.Equal(t, 101, got[0].UserId)
}

func TestGetDriverIdByUserIdAndCompanyId_UnknownReturnsZero(t *testing.T) {
	testutil.SetupTestDB(t)

	require.Zero(t, model.GetDriverIdByUserIdAndCompanyId(999, 999))
}

func TestGetUserIdByDriverId_ReturnsAssociatedUser(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriver(model.NewDriver(101, 1, "addr-a", true))
	driverRowId := model.GetDriverIdByUserIdAndCompanyId(101, 1)
	require.NotZero(t, driverRowId)

	require.Equal(t, 101, model.GetUserIdByDriverId(driverRowId))
}

func TestDeleteDriver_RemovesRow(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriver(model.NewDriver(101, 1, "addr-a", true))
	model.DeleteDriver(model.GetDriverById(101))

	require.Empty(t, model.GetAllDrivers())
}
