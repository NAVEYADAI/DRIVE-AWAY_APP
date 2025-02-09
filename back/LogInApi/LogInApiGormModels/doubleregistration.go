package LogInApiGormModels

import "back/LogInApi/config"

type DoubleRegistration struct {
	ID                    int `gorm:"column:id"`
	UserId                int `gorm:"column:userid"`
	PersonalInformationId int `gorm:"column:personalinformationid"`
}

func (receiver DoubleRegistration) TableName() string {
	return "doubleregistration"
}
func AddDoubleRegistration(userid int, personalinformationid int) {
	config.DB.Create(&DoubleRegistration{UserId: userid, PersonalInformationId: personalinformationid})
}
func GetDoubleRegistrationById(id int) DoubleRegistration {
	var doubleRegistration DoubleRegistration
	config.DB.Where("id = ?", id).First(&doubleRegistration)
	return doubleRegistration
}
func GetAllDoubleRegistration() []DoubleRegistration {
	var doubleRegistration []DoubleRegistration
	config.DB.Find(&doubleRegistration)
	return doubleRegistration
}
func UpdateDoubleRegistration(doubleRegistration DoubleRegistration) {
	config.DB.Save(&doubleRegistration)
}
func DeleteDoubleRegistration(doubleRegistration DoubleRegistration) {
	config.DB.Delete(&doubleRegistration)
}
func GetAllPersonalInformaionById(ID int) []DoubleRegistration {
	var tmp []DoubleRegistration
	config.DB.Where("userid = ?", ID).Find(&tmp)
	return tmp
}
