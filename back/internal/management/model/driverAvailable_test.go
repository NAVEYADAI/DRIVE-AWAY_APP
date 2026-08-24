package model_test

import (
	"testing"

	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestCreateDriverAvailable_And_GetAll(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriverAvailable(model.NewDriverAvailable(1, 101, true))

	all := model.GetAllDriverAvailables()
	require.Len(t, all, 1)
	require.Equal(t, 101, all[0].IdDriver)
}

func TestGetAllDriversAviableInCompany_FiltersByCompany(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriverAvailable(model.NewDriverAvailable(1, 101, true))
	model.CreateDriverAvailable(model.NewDriverAvailable(2, 202, true))

	got := model.GetAllDriversAviableInCompany(1)
	require.Len(t, got, 1)
	require.Equal(t, 101, got[0].IdDriver)
}

func TestUpdateDriverAvailableValueByUserId_Toggles(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriverAvailable(model.NewDriverAvailable(1, 101, true))

	model.UpdateDriverAvailableValueByUserId(101)
	all := model.GetAllDriverAvailables()
	require.False(t, all[0].Available)

	model.UpdateDriverAvailableValueByUserId(101)
	all = model.GetAllDriverAvailables()
	require.True(t, all[0].Available)
}

func TestDeleteDriverAvailableByDriverId_RemovesRow(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriverAvailable(model.NewDriverAvailable(1, 101, true))
	model.DeleteDriverAvailableByDriverId(101)

	require.Empty(t, model.GetAllDriverAvailables())
}
