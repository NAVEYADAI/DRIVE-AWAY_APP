package test

import (
	"back/MainController"
	"back/MainController/MainControllerGormModels"
	"fmt"
	"testing"
)

func TestDeleteDriverAvailableByDriveId(t *testing.T) {
	MainController.UpDbMainController()
	MainControllerGormModels.DeleteDriverAvailableByDriverId(11)
}
func TestUpdateDriverAvailableValueByUserId(t *testing.T) {
	MainController.UpDbMainController()
	tmpId := 12
	tmp := MainControllerGormModels.GetDriveAvailableById(tmpId)
	fmt.Println(tmp)
	MainControllerGormModels.UpdateDriverAvailableValueByUserId(tmpId)
	tmp = MainControllerGormModels.GetDriveAvailableById(tmpId)
	fmt.Println(tmp)
}
