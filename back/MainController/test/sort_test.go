package test

import (
	"back/MainController"
	"back/MainController/MainControllerGormModels"
	"testing"
)

func TestGetNumOfLastSortByCompanyId(t *testing.T) {
	MainController.UpDbMainController()
	tmp := MainControllerGormModels.GetNumOfLastSortByCompanyId(1)
	print(tmp)
}
