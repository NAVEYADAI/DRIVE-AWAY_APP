package test

import (
	"back/MainController"
	"back/MainController/MainControllerGormModels"
	"testing"
)

func TestGetDriveIdByName(t *testing.T) {
	MainController.UpDbMainController()
	id := MainControllerGormModels.GetDriveByName("tmp")
	println(id)
}
func TestGetDriveNameByID(t *testing.T) {
	MainController.UpDbMainController()
	Name := MainControllerGormModels.GetDriveNameById(1)
	println(Name)
}
