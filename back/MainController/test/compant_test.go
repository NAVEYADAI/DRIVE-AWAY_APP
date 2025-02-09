package test

import (
	"back/MainController"
	"back/MainController/MainControllerGormModels"
	"testing"
)

func TestCreateCompanyAndReturnRow(t *testing.T) {
	MainController.UpDbMainController()
	tmp := MainControllerGormModels.
		CreateCompanyWithReturnId("TestCompany", "TestCompanyId", "TestCompanyAddress")
	println(tmp.ID)
}
func TestGetIdByName(t *testing.T) {
	MainController.UpDbMainController()
	tmp := MainControllerGormModels.GetCompanyByName("yadai")
	println(tmp)
}
