package test

import (
	"back/MainController/MainControllerGormModels"
	"back/MainController/config"
	"testing"
)

func TestTryGetPermissionByUserIdAndLevel(t *testing.T) {
	config.UpDBMainController()
	tmpLevel := MainControllerGormModels.GetPermissionByUserIdAndLevel(5, 1)
	println(tmpLevel.GetId())
}
func TestTryGetPermissionByUserIdAndLevelIfWeDidntHaveRow(t *testing.T) {
	config.UpDBMainController()
	tmpLevel := MainControllerGormModels.GetPermissionByUserIdAndLevel(5, 4)
	if tmpLevel.Level != 0 {
		println(tmpLevel.GetCompanyId())
		println("")
	}
}
func TestGetCompanyIdByLevelAndUserId(t *testing.T) {
	config.UpDBMainController()

	per := MainControllerGormModels.GetCompanyIdByLevelAndUserId(5, 1)
	println(per)
}
func TestGetIdByAllIds(t *testing.T) {
	config.UpDBMainController()

	per := MainControllerGormModels.GetPermissionsByAllId(5, 1, 1)
	println(per)
}
func TestGetUserIdByCompanyIdAndLevel(t *testing.T) {
	config.UpDBMainController()

	per := MainControllerGormModels.GetUserIdByCompanyIdAndLevel(5, 1)
	println(per)
}
