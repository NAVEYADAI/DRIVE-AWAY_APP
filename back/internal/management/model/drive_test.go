package model_test

import (
	"testing"

	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestCreateDriveReturnId_ReturnsUsableId(t *testing.T) {
	testutil.SetupTestDB(t)

	id := model.CreateDriveReturnId(model.NewDrive("Morning Run", 8, 30, 1, "10,10", "20,20"))
	require.NotZero(t, id)

	got := model.GetDriveById(id)
	require.Equal(t, "Morning Run", got.Name)
	require.Equal(t, 8, got.HToStart)
	require.Equal(t, 30, got.MToStart)
}

func TestGetDriveByName_DistinguishesRows(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreateDrive(model.NewDrive("Morning Run", 8, 0, 1, "10,10", "20,20"))
	eveningId := model.CreateDriveReturnId(model.NewDrive("Evening Run", 18, 0, 1, "10,10", "20,20"))

	got := model.GetDriveByName("Evening Run")
	require.Equal(t, eveningId, got)
}

func TestGetDriveByName_UnknownNameReturnsZero(t *testing.T) {
	testutil.SetupTestDB(t)

	require.Zero(t, model.GetDriveByName("Nonexistent"))
}

func TestGetDriveNameById_UnknownIdReturnsEmptyString(t *testing.T) {
	testutil.SetupTestDB(t)

	require.Empty(t, model.GetDriveNameById(999999))
}

func TestUpdateDrive_PersistsChanges(t *testing.T) {
	testutil.SetupTestDB(t)

	id := model.CreateDriveReturnId(model.NewDrive("Morning Run", 8, 0, 1, "10,10", "20,20"))
	drive := model.GetDriveById(id)
	drive.Name = "Renamed Run"
	model.UpdateDrive(drive)

	require.Equal(t, "Renamed Run", model.GetDriveById(id).Name)
}

func TestDeleteDrive_RemovesRow(t *testing.T) {
	testutil.SetupTestDB(t)

	id := model.CreateDriveReturnId(model.NewDrive("Morning Run", 8, 0, 1, "10,10", "20,20"))
	model.DeleteDrive(model.GetDriveById(id))

	require.Empty(t, model.GetAllDrives())
}
