package LogInApiLogic

import (
	"back/LogInApi/LogInApiGormModels"
	"regexp"
)

func isUserNameTaken(newUserName string) bool {
	userNameList := LogInApiGormModels.GetAllUser()
	for _, userName := range userNameList {
		if userName.UserName == newUserName {
			return true
		}
	}
	return false
}

func isUserNameLengthValid(newUserName string) bool {
	return len(newUserName) >= 5 && len(newUserName) <= 20
}

func isUserNameFormatValid(newUserName string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", newUserName)
	return matched
}

func isUserNamePersonalInfoValid(newUserName string) bool {
	matched, _ := regexp.MatchString("\\d{4}-\\d{2}-\\d{2}", newUserName)
	return !matched
}

func isUserNameStartWithLetter(newUserName string) bool {
	return (newUserName[0] >= 'A' && newUserName[0] <= 'Z') || (newUserName[0] >= 'a' && newUserName[0] <= 'z')
}

func SignUpNewUserName(newUserName string) (bool, string) {
	if isUserNameTaken(newUserName) {
		return false, "User name already taken"
	}

	if !isUserNameLengthValid(newUserName) {
		return false, "User name must be between 5 and 20 characters"
	}

	if !isUserNameFormatValid(newUserName) {
		return false, "User name must contain only letters and numbers"
	}

	if !isUserNamePersonalInfoValid(newUserName) {
		return false, "User name must not contain personal information"
	}

	if !isUserNameStartWithLetter(newUserName) {
		return false, "User name must start with a letter"
	}

	return true, "Good user name"
}
