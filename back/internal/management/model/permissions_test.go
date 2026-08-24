package model_test

import (
	"testing"

	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestCreatePermissions_And_GetByUserId(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreatePermissions(model.NewPermissions(101, 1, 3))

	got := model.GetPermissionsByUserId(101)
	require.Equal(t, 1, got.ConpanyId)
	require.Equal(t, 3, got.Level)
}

func TestGetPermissionLevelByUserId_UnknownUserReturnsZero(t *testing.T) {
	testutil.SetupTestDB(t)

	require.Zero(t, model.GetPermissionLevelByUserId(999))
}

func TestGetCompanyIdByLevelAndUserId_MatchesBothFields(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreatePermissions(model.NewPermissions(101, 1, 3))
	model.CreatePermissions(model.NewPermissions(101, 2, 5)) // same user, different level/company

	require.Equal(t, 1, model.GetCompanyIdByLevelAndUserId(101, 3))
	require.Equal(t, 2, model.GetCompanyIdByLevelAndUserId(101, 5))
}

func TestGetUserIdByCompanyIdAndLevel_MatchesBothFields(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreatePermissions(model.NewPermissions(101, 1, 3))
	model.CreatePermissions(model.NewPermissions(202, 1, 5))

	require.Equal(t, 101, model.GetUserIdByCompanyIdAndLevel(3, 1))
	require.Equal(t, 202, model.GetUserIdByCompanyIdAndLevel(5, 1))
}

func TestGetPermissionByUserIdAndLevel_NoMatchReturnsZeroValue(t *testing.T) {
	testutil.SetupTestDB(t)

	got := model.GetPermissionByUserIdAndLevel(999, 1)
	require.Zero(t, got.ID)
	require.Zero(t, got.Level)
}

func TestDeletePermissions_RemovesRow(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreatePermissions(model.NewPermissions(101, 1, 3))
	p := model.GetPermissionsByUserId(101)
	model.DeletePermissions(p)

	require.Empty(t, model.GetAllPermissions())
}
