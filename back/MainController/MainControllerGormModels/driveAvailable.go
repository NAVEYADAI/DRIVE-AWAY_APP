package MainControllerGormModels

import (
	"back/MainController/config"
	"fmt"
)

type DriveAvailable struct {
	ID        int  `gorm:"column:id"`
	IdCompany int  `gorm:"column:idcompany"`
	IdDrive   int  `gorm:"column:iddrive"`
	Available bool `gorm:"column:available"`
}

func (receiver DriveAvailable) TableName() string {
	return "driveavailable"
}
func NewDriveAvailable(idCompany, idDrive int, available bool) DriveAvailable {
	return DriveAvailable{IdCompany: idCompany, IdDrive: idDrive, Available: available}
}
func (receiver DriveAvailable) GetId() int {
	return receiver.ID
}
func (receiver DriveAvailable) GetIdCompany() int {
	return receiver.IdCompany
}
func (receiver DriveAvailable) GetIdDrive() int {
	return receiver.IdDrive
}
func (receiver DriveAvailable) GetAvailable() bool {
	return receiver.Available
}
func (receiver DriveAvailable) SetIdCompany(idCompany int) {
	receiver.IdCompany = idCompany
}
func (receiver DriveAvailable) SetIdDrive(idDrive int) {
	receiver.IdDrive = idDrive
}
func (receiver DriveAvailable) SetAvailable() {
	receiver.Available = !receiver.Available
}

// Gorm functions
func CreateDriveAvailable(receiver DriveAvailable) {
	config.DB.Create(&receiver)
}
func GetAllDriveAvailables() []DriveAvailable {
	var driveAvailables []DriveAvailable
	config.DB.Find(&driveAvailables)
	return driveAvailables
}
func GetDriveAvailableById(id int) DriveAvailable {
	var driveAvailable DriveAvailable
	config.DB.Where("id = ?", id).Find(&driveAvailable)
	return driveAvailable
}
func GetDriveAviablesByCompanyId(companyId int) []DriveAvailable {
	var driveAvailable []DriveAvailable
	config.DB.Where("idcompany = ?", companyId).Find(&driveAvailable)
	return driveAvailable
}
func UpdateDriveAvailable(receiver DriveAvailable) {
	config.DB.Save(&receiver)
}
func DeleteDriveAvailable(receiver DriveAvailable) {
	config.DB.Delete(&receiver)
}
func DeleteDriveAvailableByDriveId(driveId int) {
	config.DB.Where("iddrive = ?", driveId).Delete(DriveAvailable{})
}
func UpdateDriveAvailableValueByDriveId(driveId int) {
	// if Available == true set to false and if him false set to true
	var driveAvailable DriveAvailable

	// מצא את הרשומה עם driveId המסופק
	result := config.DB.Where("iddrive = ?", driveId).First(&driveAvailable)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	// עדכן את הערך של Available
	newAvailableValue := !driveAvailable.Available
	result = config.DB.Model(&driveAvailable).Update("available", newAvailableValue)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

}
func GetAvailableDriveByDriveId(driveId int) DriveAvailable {
	var tmp DriveAvailable
	config.DB.Where("iddrive", driveId).Find(&tmp)
	return tmp
}
