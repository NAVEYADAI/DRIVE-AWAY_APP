package MainControllerGormModels

import "back/MainController/config"

type Sort struct {
	ID        int    `gorm:"column:id"`
	CompanyId int    `gorm:"column:companyid"`
	Name      string `gorm:"column:name"`
}

func (receiver Sort) TableName() string {
	return "sort"
}
func CreateSort(companyId int, Name string) {
	tmpSort := Sort{CompanyId: companyId, Name: Name}
	config.DB.Create(&tmpSort)
}
func GetLastSortByCompanyId(companyId int) Sort {
	var tmp Sort
	config.DB.Where("companyid = ?", companyId).Last(&tmp)
	return tmp
}
func GetNumOfLastSortByCompanyId(companyId int) int {
	var tmp Sort
	config.DB.Where("companyid = ?", companyId).Last(&tmp)
	return tmp.ID
}
