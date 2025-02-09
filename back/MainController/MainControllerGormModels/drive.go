package MainControllerGormModels

import "back/MainController/config"

type Drive struct {
	ID        int    `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	HToStart  int    `gorm:"column:htostart"`
	MToStart  int    `gorm:"column:mtostart"`
	GpsStart  string `gorm:"column:gpsstart"`
	GpsEnd    string `gorm:"column:gpsend"`
	CompanyId int    `gorm:"column:companyid"`
}

func (receiver Drive) TableName() string {
	return "drive"
}
func NewDrive(name string, hToStart, mToStart, companyId int, gpsStart, gpsEnd string) Drive {
	return Drive{Name: name, HToStart: hToStart, MToStart: mToStart, GpsStart: gpsStart, GpsEnd: gpsEnd, CompanyId: companyId}
}
func (receiver Drive) GetId() int {
	return receiver.ID
}
func (receiver Drive) GetName() string {
	return receiver.Name
}

func (receiver Drive) GetHToStart() int {
	return receiver.HToStart
}
func (receiver Drive) GetMToStart() int {
	return receiver.MToStart
}
func (receiver Drive) GetGPSEnd() string {
	return receiver.GpsEnd
}
func (receiver Drive) GetGPSStart() string {
	return receiver.GpsStart
}
func (receiver Drive) SetName(name string) {
	receiver.Name = name
}
func (receiver Drive) SetHToStart(hToStart int) {
	receiver.HToStart = hToStart
}
func (receiver Drive) SetMToStart(mToStart int) {
	receiver.MToStart = mToStart
}
func (receiver Drive) SetGPSEnd(gpsEnd string) {
	receiver.GpsEnd = gpsEnd
}
func (receiver Drive) SetGPSStart(gpsStart string) {
	receiver.GpsStart = gpsStart
}

// Gorm functions
func CreateDrive(receiver Drive) {
	config.DB.Create(&receiver)
}
func CreateDriveReturnId(receiver Drive) int {
	config.DB.Create(&receiver)
	return receiver.ID
}
func GetAllDrives() []Drive {
	var drives []Drive
	config.DB.Find(&drives)
	return drives
}
func GetDriveById(id int) Drive {
	var drive Drive
	config.DB.Where("id = ?", id).Find(&drive)
	return drive
}
func GetDriveByName(name string) int {
	var drive Drive
	config.DB.Where("name", name).Find(&drive)
	return drive.ID

}
func UpdateDrive(receiver Drive) {
	config.DB.Save(&receiver)
}
func DeleteDrive(receiver Drive) {
	config.DB.Delete(&receiver)
}

func GetDriveNameById(id int) string {
	var drive Drive
	config.DB.Where("id = ?", id).Find(&drive)
	return drive.Name
}
