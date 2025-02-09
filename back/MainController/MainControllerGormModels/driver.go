package MainControllerGormModels

import "back/MainController/config"

type Driver struct {
	ID         int    `gorm:"column:id"`
	UserId     int    `gorm:"column:userid"`
	CompanyId  int    `gorm:"column:companyid"`
	Address    string `gorm:"column:address"`
	IsEmployee bool   `gorm:"column:isemployee"`
}

func (receiver Driver) TableName() string {
	return "driver"
}
func NewDriver(userId, companyId int, address string, isEmployee bool) Driver {
	return Driver{UserId: userId, CompanyId: companyId, Address: address, IsEmployee: isEmployee}
}
func (receiver Driver) GetId() int {
	return receiver.ID
}
func (receiver Driver) GetAddress() string {
	return receiver.Address
}
func (receiver Driver) GetIsEmployee() bool {
	return receiver.IsEmployee
}
func (receiver Driver) SetAddress(address string) {
	receiver.Address = address
}
func (receiver Driver) SetIsEmployee() {
	receiver.IsEmployee = !receiver.IsEmployee
}

// Gorm functions
func CreateDriver(receiver Driver) {
	config.DB.Create(&receiver)
}
func GetAllDrivers() []Driver {
	var drivers []Driver
	config.DB.Find(&drivers)
	return drivers
}
func GetAllDriversInCompany(CompanyId int) []Driver {
	var drivers []Driver
	config.DB.Where("companyid", CompanyId).Find(&drivers)
	return drivers
}
func GetDriverById(id int) Driver {
	var driver Driver
	config.DB.Where("userid = ?", id).Find(&driver)
	return driver
}
func UpdateDriver(receiver Driver) {
	config.DB.Save(&receiver)
}
func DeleteDriver(receiver Driver) {
	config.DB.Delete(&receiver)
}
func GetDriverIdByUserIdAndCompanyId(userId, companyId int) int {
	var driver Driver
	config.DB.Where("userid = ? AND companyid = ?", userId, companyId).Find(&driver)
	return driver.ID
}
func GetUserIdByDriverId(driverId int) int {
	var driver Driver
	config.DB.Where("id = ?", driverId).Find(&driver)
	return driver.UserId
}
