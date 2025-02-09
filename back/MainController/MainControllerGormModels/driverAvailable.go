package MainControllerGormModels

import (
	"back/MainController/config"
	"fmt"
)

type DriverAvailable struct {
	ID        int  `gorm:"column:id"`
	IdCompany int  `gorm:"column:idcompany"`
	IdDriver  int  `gorm:"column:iddriver"`
	Available bool `gorm:"column:available"`
}

func (receiver DriverAvailable) TableName() string {
	return "driveravailable"
}
func NewDriverAvailable(idCompany, idDriver int, available bool) DriverAvailable {
	return DriverAvailable{IdCompany: idCompany, IdDriver: idDriver, Available: available}
}
func (receiver DriverAvailable) GetId() int {
	return receiver.ID
}
func (receiver DriverAvailable) GetIdCompany() int {
	return receiver.IdCompany
}
func (receiver DriverAvailable) GetIdDriver() int {
	return receiver.IdDriver
}
func (receiver DriverAvailable) GetAvailable() bool {
	return receiver.Available
}
func (receiver DriverAvailable) SetIdCompany(idCompany int) {
	receiver.IdCompany = idCompany
}
func (receiver DriverAvailable) SetIdDriver(idDriver int) {
	receiver.IdDriver = idDriver
}
func (receiver DriverAvailable) SetAvailable() {
	receiver.Available = !receiver.Available
}

// Gorm functions
func CreateDriverAvailable(receiver DriverAvailable) {
	config.DB.Create(&receiver)
}
func GetAllDriverAvailables() []DriverAvailable {
	var driverAvailables []DriverAvailable
	config.DB.Find(&driverAvailables)
	return driverAvailables
}
func GetDriverAvailableById(id int) DriverAvailable {
	var driverAvailable DriverAvailable
	config.DB.Where("ID = ?", id).Find(&driverAvailable)
	return driverAvailable
}
func GetAllDriversAviableInCompany(CompanyId int) []DriverAvailable {
	var drivers []DriverAvailable
	config.DB.Where("idcompany", CompanyId).Find(&drivers)
	return drivers
}

func UpdateDriverAvailable(receiver DriverAvailable) {
	config.DB.Save(&receiver)
}
func DeleteDriverAvailableByDriverId(driverId int) {
	config.DB.Where("iddriver = ?", driverId).Delete(&DriverAvailable{})
}
func DeleteDriverAvailable(receiver DriverAvailable) {
	config.DB.Delete(&receiver)
}
func UpdateDriverAvailableValueByUserId(driverId int) {
	var driverAvailable DriverAvailable

	// מצא את הרשומה עם driveId המסופק
	result := config.DB.Where("iddriver = ?", driverId).First(&driverAvailable)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	// עדכן את הערך של Available
	newAvailableValue := !driverAvailable.Available
	result = config.DB.Model(&driverAvailable).Update("available", newAvailableValue)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
