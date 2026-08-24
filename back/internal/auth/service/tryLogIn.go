package service

import (
	"back/internal/auth/model"
	"time"
)

func TryLogInIpAndMAcAddress(IP string, MacAddress string, Date time.Time) (bool, string) {
	ListTryLogIn := model.GetTryLogInByIpAndMacAddress(IP, MacAddress, Date)
	DidNotSucceed := 0
	for _, tryLogIn := range ListTryLogIn {
		if !tryLogIn.Connect {
			DidNotSucceed++
		}
	}
	if DidNotSucceed > 5 {
		return false, "Too many attempts"
	}
	return true, ""
}
