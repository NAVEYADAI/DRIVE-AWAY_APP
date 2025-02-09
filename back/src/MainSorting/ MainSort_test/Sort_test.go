package main

import (
	"back/MainController"
	"back/MainController/MainControllerGormModels"
	"fmt"
	"testing"
)

func TestGetLastSortByCompanyId(t *testing.T) {
	MainController.UpDbMainController()
	tmp := MainControllerGormModels.GetLastSortByCompanyId(1)
	fmt.Println(tmp)
}
func TestGetNumOfLastSortByCompanyId(t *testing.T) {
	MainController.UpDbMainController()
	tmp := MainControllerGormModels.GetNumOfLastSortByCompanyId(1)
	fmt.Println(tmp)
}
func TestCreateSort(t *testing.T) {
	MainController.UpDbMainController()
	MainControllerGormModels.CreateSort(1, "testCreated")
	tmp := MainControllerGormModels.GetLastSortByCompanyId(1)
	fmt.Println(tmp)
}
