package model_test

import (
	"testing"

	"back/internal/auth/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDoubleRegistration_AndGetById(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddPersonalInformation("Alice", "Smith", "123456789", userId)
	pi := model.GetAllPersonalInformation()
	require.Len(t, pi, 1)

	model.AddDoubleRegistration(userId, pi[0].ID)

	all := model.GetAllDoubleRegistration()
	require.Len(t, all, 1)
	assert.Equal(t, userId, all[0].UserId)
	assert.Equal(t, pi[0].ID, all[0].PersonalInformationId)

	got := model.GetDoubleRegistrationById(all[0].ID)
	assert.Equal(t, userId, got.UserId)
}

func TestGetAllPersonalInformaionById(t *testing.T) {
	testutil.SetupTestDB(t)

	alice := model.AddUser("alice")
	bob := model.AddUser("bob")
	model.AddDoubleRegistration(alice, 1)
	model.AddDoubleRegistration(alice, 2)
	model.AddDoubleRegistration(bob, 3)

	got := model.GetAllPersonalInformaionById(alice)
	assert.Len(t, got, 2)
}

func TestUpdateAndDeleteDoubleRegistration(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	model.AddDoubleRegistration(userId, 1)
	all := model.GetAllDoubleRegistration()
	require.Len(t, all, 1)

	entry := all[0]
	entry.PersonalInformationId = 2
	model.UpdateDoubleRegistration(entry)
	assert.Equal(t, 2, model.GetDoubleRegistrationById(entry.ID).PersonalInformationId)

	model.DeleteDoubleRegistration(entry)
	assert.Empty(t, model.GetAllDoubleRegistration())
}
