package service_test

import (
	"testing"

	"back/internal/auth/service"

	"github.com/stretchr/testify/assert"
)

func TestIsSecure_ValidPassword(t *testing.T) {
	good, message := service.IsSecure("Abcdef1!")
	assert.True(t, good)
	assert.Equal(t, "Password is secure", message)
}

func TestIsSecure_TooShort(t *testing.T) {
	good, message := service.IsSecure("Ab1!")
	assert.False(t, good)
	assert.Equal(t, "Password must be at least 8 characters long", message)
}

func TestIsSecure_NoUppercase(t *testing.T) {
	good, message := service.IsSecure("abcdefg1!")
	assert.False(t, good)
	assert.Equal(t, "Password must contain at least one uppercase letter", message)
}

func TestIsSecure_NoLowercase(t *testing.T) {
	good, message := service.IsSecure("ABCDEFG1!")
	assert.False(t, good)
	assert.Equal(t, "Password must contain at least one lowercase letter", message)
}

func TestIsSecure_NoDigit(t *testing.T) {
	good, message := service.IsSecure("Abcdefgh!")
	assert.False(t, good)
	assert.Equal(t, "Password must contain at least one digit", message)
}

func TestIsSecure_NoSpecialChar(t *testing.T) {
	good, message := service.IsSecure("Abcdefg1")
	assert.False(t, good)
	assert.Equal(t, "Password must contain at least one special character", message)
}

func TestIsSecure_EmptyPassword(t *testing.T) {
	good, _ := service.IsSecure("")
	assert.False(t, good)
}
