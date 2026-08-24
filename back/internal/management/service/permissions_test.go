package service_test

import (
	"testing"

	"back/internal/management/model"
	"back/internal/management/service"
	"back/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestReturnPermissionForSingUp_AdminLevelsMapToOne(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreatePermissions(model.NewPermissions(101, 1, 1))
	require.Equal(t, 1, service.ReturnPermissionForSingUp(101, 1))

	model.CreatePermissions(model.NewPermissions(102, 1, 2))
	require.Equal(t, 1, service.ReturnPermissionForSingUp(102, 2))
}

func TestReturnPermissionForSingUp_ManagerLevelsMapToTwo(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreatePermissions(model.NewPermissions(101, 1, 3))
	require.Equal(t, 2, service.ReturnPermissionForSingUp(101, 3))

	model.CreatePermissions(model.NewPermissions(102, 1, 4))
	require.Equal(t, 2, service.ReturnPermissionForSingUp(102, 4))
}

func TestReturnPermissionForSingUp_OtherKnownLevelReturnsZero(t *testing.T) {
	testutil.SetupTestDB(t)

	model.CreatePermissions(model.NewPermissions(101, 1, 5))
	require.Equal(t, 0, service.ReturnPermissionForSingUp(101, 5))
}

func TestReturnPermissionForSingUp_NoPermissionRowReturnsNegativeOne(t *testing.T) {
	testutil.SetupTestDB(t)

	require.Equal(t, -1, service.ReturnPermissionForSingUp(999, 1))
}
