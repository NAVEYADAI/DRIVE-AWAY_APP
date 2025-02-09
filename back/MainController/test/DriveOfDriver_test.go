package test

import (
	"back/MainController"
	"back/MainController/MainControllerGormModels"
	"fmt"
	"testing"
)

func TestGetAllDrivesByDriverID(t *testing.T) {
	MainController.UpDbMainController()
	tmp := MainControllerGormModels.GetAllDrivesByDriverID(12, 7)
	fmt.Println(tmp)
}
