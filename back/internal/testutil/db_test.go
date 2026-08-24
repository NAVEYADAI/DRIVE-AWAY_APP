package testutil

import (
	"testing"

	authmodel "back/internal/auth/model"

	"github.com/stretchr/testify/require"
)

// TestSetupTestDB_IsIsolatedFromDev is a smoke test proving SetupTestDB
// connects to db_test (not the dev/production DB) and starts each run clean.
func TestSetupTestDB_IsIsolatedFromDev(t *testing.T) {
	db := SetupTestDB(t)

	var count int64
	require.NoError(t, db.Model(&authmodel.Users{}).Count(&count).Error)
	require.Equal(t, int64(0), count, "expected a freshly truncated users table")

	require.NoError(t, db.Create(&authmodel.Users{UserName: "smoke-test-user"}).Error)
	require.NoError(t, db.Model(&authmodel.Users{}).Count(&count).Error)
	require.Equal(t, int64(1), count)
}
