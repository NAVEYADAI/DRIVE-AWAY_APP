package model_test

import (
	"testing"

	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestCreateDriveAvailable_And_GetAll(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriveAvailable(model.NewDriveAvailable(1, 42, true))

	all := model.GetAllDriveAvailables()
	require.Len(t, all, 1)
	require.Equal(t, 42, all[0].IdDrive)
	require.True(t, all[0].Available)
}

func TestGetDriveAviablesByCompanyId_FiltersByCompany(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriveAvailable(model.NewDriveAvailable(1, 10, true))
	model.CreateDriveAvailable(model.NewDriveAvailable(2, 20, true))

	got := model.GetDriveAviablesByCompanyId(1)
	require.Len(t, got, 1)
	require.Equal(t, 10, got[0].IdDrive)
}

func TestUpdateDriveAvailableValueByDriveId_Toggles(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriveAvailable(model.NewDriveAvailable(1, 42, true))

	model.UpdateDriveAvailableValueByDriveId(42)
	require.False(t, model.GetAvailableDriveByDriveId(42).Available, "first toggle should flip true -> false")

	model.UpdateDriveAvailableValueByDriveId(42)
	require.True(t, model.GetAvailableDriveByDriveId(42).Available, "second toggle should flip back false -> true")
}

func TestDeleteDriveAvailableByDriveId_RemovesRow(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDriveAvailable(model.NewDriveAvailable(1, 42, true))
	model.DeleteDriveAvailableByDriveId(42)

	require.Empty(t, model.GetAllDriveAvailables())
}
