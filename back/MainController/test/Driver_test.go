package test

import (
	"back/MainController"
	"back/MainController/MainControllerGormModels"
	"fmt"
	"testing"
)

func TestGetDriverByName(t *testing.T) {
	MainController.UpDbMainController()
	tmp := MainControllerGormModels.GetAllDriversInCompany(1)
	fmt.Println(tmp)
}
func TestGetDriverIdByUserIdAndCompanyId(t *testing.T) {
	MainController.UpDbMainController()
	tmp := MainControllerGormModels.GetDriverIdByUserIdAndCompanyId(10, 1)
	fmt.Println(tmp)
}
