package model_test

import (
	"testing"

	"back/internal/auth/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddPersonalInformation_AndGetById(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPersonalInformation("Alice", "Smith", "123456789", userId)

	all := model.GetAllPersonalInformation()
	require.Len(t, all, 1)
	assert.Equal(t, "Alice", all[0].FName)
	assert.Equal(t, "Smith", all[0].LName)
	assert.Equal(t, "123456789", all[0].IdentityCard)

	got := model.GetPersonalInformationById(all[0].ID)
	assert.Equal(t, "Alice", got.FName)
}

func TestGetFNameByUserID(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPersonalInformation("Alice", "Smith", "123456789", userId)

	assert.Equal(t, "Alice", model.GetFNameByUserID(userId))
}

func TestGetFNameByUserID_NoRecord(t *testing.T) {
	testutil.SetupTestDB(t)
	assert.Equal(t, "", model.GetFNameByUserID(999999))
}

func TestGetPersonalInformationByIDCard(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPersonalInformation("Alice", "Smith", "123456789", userId)

	got := model.GetPersonalInformationByIDCard("123456789")
	assert.Equal(t, userId, got.UserId)
}

// TestGetUserIdFromID is a regression test: the query compared the
// identitycard text column directly against a bare int parameter, which
// the Postgres driver can't encode ("unable to encode 123456789 into text
// format for text") — the query errored on every call and GetUserIdFromID
// always returned -1, even for a card number that existed.
func TestGetUserIdFromID(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPersonalInformation("Alice", "Smith", "123456789", userId)

	assert.Equal(t, userId, model.GetUserIdFromID(123456789))
}

func TestGetUserIdFromID_NotFound(t *testing.T) {
	testutil.SetupTestDB(t)
	assert.Equal(t, -1, model.GetUserIdFromID(999))
}

func TestUpdateAndDeletePersonalInformation(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPersonalInformation("Alice", "Smith", "123456789", userId)
	all := model.GetAllPersonalInformation()
	require.Len(t, all, 1)

	entry := all[0]
	entry.LName = "Jones"
	model.UpdatePersonalInformation(entry)
	assert.Equal(t, "Jones", model.GetPersonalInformationById(entry.ID).LName)

	model.DeletePersonalInformation(entry)
	assert.Empty(t, model.GetAllPersonalInformation())
}
