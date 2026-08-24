package model_test

import (
	"testing"

	"back/internal/auth/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddUser_AndGetUserById(t *testing.T) {
	testutil.SetupTestDB(t)

	id := model.AddUser("alice")
	require.NotZero(t, id)

	got := model.GetUserById(id)
	assert.Equal(t, "alice", got.UserName)
}

func TestGetUserById_NotFound(t *testing.T) {
	testutil.SetupTestDB(t)

	got := model.GetUserById(999999)
	assert.Equal(t, -1, got.ID)
}

func TestGetAllUser(t *testing.T) {
	testutil.SetupTestDB(t)

	model.AddUser("alice")
	model.AddUser("bob")

	users := model.GetAllUser()
	require.Len(t, users, 2)
	names := []string{users[0].UserName, users[1].UserName}
	assert.ElementsMatch(t, []string{"alice", "bob"}, names)
}

func TestUpdateUser(t *testing.T) {
	testutil.SetupTestDB(t)

	id := model.AddUser("alice")
	model.UpdateUser("alice2", id)

	got := model.GetUserById(id)
	assert.Equal(t, "alice2", got.UserName)
}

func TestDeleteUser(t *testing.T) {
	testutil.SetupTestDB(t)

	id := model.AddUser("alice")
	model.DeleteUser(model.Users{ID: id})

	got := model.GetUserById(id)
	assert.Equal(t, -1, got.ID)
}

func TestGetUserByUserName_FindsCorrectUserAmongMany(t *testing.T) {
	testutil.SetupTestDB(t)

	model.AddUser("alice")
	bobId := model.AddUser("bob")
	model.AddUser("carol")

	assert.Equal(t, bobId, model.GetUserByUserName("bob"))
}

// TestGetUserByUserName_ReturnsMinusOneWhenNotFound is a regression test: the
// function used to return the Go zero-value 0 for "not found" (from an
// unpopulated Users struct), but every caller (auth.Login, handler.LogInUser)
// checks `UserId == -1` to detect a missing user — matching the convention
// GetUserById already uses in this same file. Because 0 != -1, the "user not
// found" branch never fired; login instead fell through to comparing an
// empty stored password against the submitted one for user id 0, returning
// the wrong HTTP status/message for a nonexistent username.
func TestGetUserByUserName_ReturnsMinusOneWhenNotFound(t *testing.T) {
	testutil.SetupTestDB(t)

	got := model.GetUserByUserName("does-not-exist")
	assert.Equal(t, -1, got, "must match the -1 'not found' convention used by GetUserById and expected by callers")
}
