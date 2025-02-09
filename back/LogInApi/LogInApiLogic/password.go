package LogInApiLogic

import "unicode"

func isLengthValid(password string) bool {
	return len(password) >= 8
}

func containsUppercase(password string) bool {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func containsLowercase(password string) bool {
	for _, char := range password {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func containsDigit(password string) bool {
	for _, char := range password {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func containsSpecialChar(password string) bool {
	specialChars := "!@#$%^&*()-_+=[]{}|;:'\",.<>?/"
	for _, char := range password {
		for _, special := range specialChars {
			if char == special {
				return true
			}
		}
	}
	return false
}

func IsSecure(password string) (bool, string) {
	if !isLengthValid(password) {
		return false, "Password must be at least 8 characters long"
	}

	if !containsUppercase(password) {
		return false, "Password must contain at least one uppercase letter"
	}

	if !containsLowercase(password) {
		return false, "Password must contain at least one lowercase letter"
	}

	if !containsDigit(password) {
		return false, "Password must contain at least one digit"
	}

	if !containsSpecialChar(password) {
		return false, "Password must contain at least one special character"
	}

	return true, "Password is secure"
}
