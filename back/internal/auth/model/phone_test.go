package model_test

import (
	"testing"

	"back/internal/auth/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddPhone_AndGetters(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPhone("0500000000", userId)

	byUser := model.GetPhoneByUserId(userId)
	assert.Equal(t, "0500000000", byUser.Phone)

	byId := model.GetPhoneById(byUser.ID)
	assert.Equal(t, "0500000000", byId.Phone)

	all := model.GetAllPhone()
	require.Len(t, all, 1)
}

func TestUpdateAndDeletePhone(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPhone("0500000000", userId)
	p := model.GetPhoneByUserId(userId)

	p.Phone = "0511111111"
	model.UpdatePhone(p)
	assert.Equal(t, "0511111111", model.GetPhoneById(p.ID).Phone)

	model.DeletePhone(p)
	assert.Empty(t, model.GetAllPhone())
}

func TestGetUserIdFromPhone(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPhone("0500000000", userId)

	assert.Equal(t, userId, model.GetUserIdFromPhone("0500000000"))
}

func TestGetUserIdFromPhone_NotFound(t *testing.T) {
	testutil.SetupTestDB(t)

	assert.Equal(t, -1, model.GetUserIdFromPhone("0599999999"))
}
