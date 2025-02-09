package test

import (
	"back/MainController"
	"back/MainController/MainControllerGormModels"
	"fmt"
	"testing"
)

func TestDeleteById(t *testing.T) {
	MainController.UpDbMainController()
	MainControllerGormModels.DeleteDriveAvailableByDriveId(1)
}
func TestAddDeiveAviable(t *testing.T) {
	MainController.UpDbMainController()
	MainControllerGormModels.CreateDriveAvailable(MainControllerGormModels.NewDriveAvailable(1, 3, true))
}
func TestDeleteByDriverId(t *testing.T) {
	MainController.UpDbMainController()
	MainControllerGormModels.DeleteDriveAvailableByDriveId(3)
}
func TestGetDriveAviablesByCompanyId(t *testing.T) {
	MainController.UpDbMainController()
	tmp := MainControllerGormModels.GetDriveAviablesByCompanyId(1)
	fmt.Println(tmp)
}
func TestUpdateAvailableValueByDriveId(t *testing.T) {
	MainController.UpDbMainController()
	tmpId := 6
	tmp := MainControllerGormModels.GetAvailableDriveByDriveId(tmpId)
	fmt.Println(tmp)
	MainControllerGormModels.UpdateDriveAvailableValueByDriveId(tmpId)
	tmp = MainControllerGormModels.GetAvailableDriveByDriveId(tmpId)
	fmt.Println(tmp)

}
