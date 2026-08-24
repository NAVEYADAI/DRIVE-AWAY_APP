package service_test

import (
	"testing"

	"back/internal/auth/model"
	"back/internal/auth/service"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddPersonalInformation_FirstRegistration_NoDoubleRegistrationRecorded(t *testing.T) {
	testutil.SetupTestDB(t)

	userId := model.AddUser("alice")
	service.AddPersonalInformation("Alice", "Smith", "123456789", userId)

	pi := model.GetAllPersonalInformation()
	require.Len(t, pi, 1)
	assert.Empty(t, model.GetAllDoubleRegistration(), "first-time registration of an ID card must not be flagged as a double registration")
}

func TestAddPersonalInformation_DuplicateIdCard_FlaggedAsDoubleRegistration(t *testing.T) {
	testutil.SetupTestDB(t)

	firstUser := model.AddUser("alice")
	service.AddPersonalInformation("Alice", "Smith", "123456789", firstUser)

	secondUser := model.AddUser("mallory")
	service.AddPersonalInformation("Mallory", "Evil", "123456789", secondUser)

	doubles := model.GetAllDoubleRegistration()
	require.Len(t, doubles, 1, "reusing an identity card for a second user account must be flagged")
	assert.Equal(t, secondUser, doubles[0].UserId)

	pi := model.GetAllPersonalInformation()
	assert.Len(t, pi, 2, "the second registration is still recorded, just flagged")
}
