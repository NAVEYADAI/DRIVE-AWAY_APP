package service_test

import (
	"strings"
	"testing"

	"back/internal/auth/model"
	"back/internal/auth/service"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestSignUpNewUserName_Valid(t *testing.T) {
	testutil.SetupTestDB(t)

	good, message := service.SignUpNewUserName("Alice123")
	assert.True(t, good)
	assert.Equal(t, "Good user name", message)
}

func TestSignUpNewUserName_AlreadyTaken(t *testing.T) {
	testutil.SetupTestDB(t)
	model.AddUser("Alice123")

	good, message := service.SignUpNewUserName("Alice123")
	assert.False(t, good)
	assert.Equal(t, "User name already taken", message)
}

func TestSignUpNewUserName_TooShort(t *testing.T) {
	testutil.SetupTestDB(t)

	good, message := service.SignUpNewUserName("ab1")
	assert.False(t, good)
	assert.Equal(t, "User name must be between 5 and 20 characters", message)
}

func TestSignUpNewUserName_TooLong(t *testing.T) {
	testutil.SetupTestDB(t)

	good, _ := service.SignUpNewUserName(strings.Repeat("a", 21))
	assert.False(t, good)
}

func TestSignUpNewUserName_InvalidCharacters(t *testing.T) {
	testutil.SetupTestDB(t)

	good, message := service.SignUpNewUserName("alice_123")
	assert.False(t, good)
	assert.Equal(t, "User name must contain only letters and numbers", message)
}

func TestSignUpNewUserName_LooksLikeADate(t *testing.T) {
	testutil.SetupTestDB(t)

	// isUserNamePersonalInfoValid rejects anything matching \d{4}-\d{2}-\d{2}
	// (guards against accidentally using a birthdate as a username) — but
	// dashes already fail the "letters and numbers only" check first, so
	// this path is effectively unreachable through the public entry point.
	// Documented here rather than silently assumed.
	good, message := service.SignUpNewUserName("2024-01-01")
	assert.False(t, good)
	assert.Equal(t, "User name must contain only letters and numbers", message)
}

func TestSignUpNewUserName_MustStartWithLetter(t *testing.T) {
	testutil.SetupTestDB(t)

	good, message := service.SignUpNewUserName("1alice")
	assert.False(t, good)
	assert.Equal(t, "User name must start with a letter", message)
}
