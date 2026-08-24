package model_test

import (
	"testing"

	"back/internal/auth/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddPassword_AndGetPasswordById(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPassword("hunter2", userId)

	all := model.GetAllPassword()
	require.Len(t, all, 1)
	assert.Equal(t, "hunter2", all[0].Pass)

	got := model.GetPasswordById(all[0].ID)
	assert.Equal(t, "hunter2", got.Pass)
}

func TestGetLastPasswordByUserId_ReturnsMostRecent(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPassword("old-password", userId)
	model.AddPassword("new-password", userId)

	got := model.GetLastPasswordByUserId(userId)
	assert.Equal(t, "new-password", got.Pass, "changing a password should make the new one the one checked at login")
}

func TestGetLastPasswordByUserId_IgnoresOtherUsers(t *testing.T) {
	testutil.SetupTestDB(t)

	alice := model.AddUser("alice")
	bob := model.AddUser("bob")
	model.AddPassword("alice-pw", alice)
	model.AddPassword("bob-pw", bob)

	got := model.GetLastPasswordByUserId(alice)
	assert.Equal(t, "alice-pw", got.Pass)
}

func TestUpdatePassword(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPassword("original", userId)
	all := model.GetAllPassword()
	require.Len(t, all, 1)

	entry := all[0]
	entry.Pass = "changed"
	model.UpdatePassword(entry)

	got := model.GetPasswordById(entry.ID)
	assert.Equal(t, "changed", got.Pass)
}

func TestDeletePassword(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPassword("to-delete", userId)
	all := model.GetAllPassword()
	require.Len(t, all, 1)

	model.DeletePassword(all[0])
	assert.Empty(t, model.GetAllPassword())
}
