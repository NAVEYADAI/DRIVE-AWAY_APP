package model_test

import (
	"testing"

	"back/internal/auth/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddMail_AndGetMailByUserId(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddMail("alice@example.com", userId)

	got := model.GetMailByUserId(userId)
	assert.Equal(t, "alice@example.com", got.Mail)
}

func TestGetAllMail(t *testing.T) {
	testutil.SetupTestDB(t)

	a := model.AddUser("alice")
	b := model.AddUser("bob")
	model.AddMail("alice@example.com", a)
	model.AddMail("bob@example.com", b)

	all := model.GetAllMail()
	require.Len(t, all, 2)
}

func TestUpdateAndDeleteMail(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddMail("old@example.com", userId)
	m := model.GetMailByUserId(userId)

	m.Mail = "new@example.com"
	model.UpdateMail(m)
	assert.Equal(t, "new@example.com", model.GetMailById(m.ID).Mail)

	model.DeleteMail(m)
	assert.Empty(t, model.GetAllMail())
}

func TestGetUserIdFromMail(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddMail("alice@example.com", userId)

	assert.Equal(t, userId, model.GetUserIdFromMail("alice@example.com"))
}

func TestGetUserIdFromMail_NotFound(t *testing.T) {
	testutil.SetupTestDB(t)

	assert.Equal(t, -1, model.GetUserIdFromMail("nobody@example.com"))
}
