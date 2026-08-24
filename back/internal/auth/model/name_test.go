package model_test

import (
	"testing"

	"back/internal/auth/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddName_AndGetNameById(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddName("old-alice", userId)

	all := model.GetAllName()
	require.Len(t, all, 1)
	assert.Equal(t, "old-alice", all[0].Name)

	got := model.GetNameById(all[0].ID)
	assert.Equal(t, "old-alice", got.Name)
}

// TestGetUserListById_ReturnsUsernameHistory is a regression test: the query
// used the bare, unquoted identifier `idUser`, which Postgres folds to
// lowercase (`iduser`) — a column that doesn't exist, since the actual
// column (per the gorm:"column:idUser" tag) is mixed-case and needs
// double-quoting to be addressed correctly. Every call failed with
// `column "iduser" does not exist` and silently returned an empty slice.
func TestGetUserListById_ReturnsUsernameHistory(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddName("alice-v1", userId)
	model.AddName("alice-v2", userId)

	history := model.GetUserListById(userId)
	require.Len(t, history, 2)
}

func TestUpdateAndDeleteName(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddName("alice-old", userId)
	all := model.GetAllName()
	require.Len(t, all, 1)

	entry := all[0]
	entry.Name = "alice-renamed"
	model.UpdateName(entry)
	assert.Equal(t, "alice-renamed", model.GetNameById(entry.ID).Name)

	model.DeleteName(entry)
	assert.Empty(t, model.GetAllName())
}
