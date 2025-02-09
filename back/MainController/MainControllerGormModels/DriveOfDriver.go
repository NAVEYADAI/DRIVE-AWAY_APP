package MainControllerGormModels

import "back/MainController/config"

type DriveOfDriver struct {
	ID       int `gorm:"column:id"`
	SortId   int `gorm:"column:sortid"`
	DriveId  int `gorm:"column:driveid"`
	DriverId int `gorm:"column:driverid"`
}

func (receiver DriveOfDriver) TableName() string {
	return "driveofdriver"
}

func CreateDriveOfDriver(sortId, DriveId, DriverId int) {
	tmp := DriveOfDriver{SortId: sortId, DriveId: DriveId, DriverId: DriverId}
	config.DB.Create(&tmp)
}
func GetAllDrivesByDriverID(driverId, sortId int) []DriveOfDriver {
	var tmp []DriveOfDriver
	config.DB.Where("driverid = ? AND sortid = ?", driverId, sortId).Find(&tmp)
	return tmp
}
