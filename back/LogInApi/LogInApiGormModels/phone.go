package LogInApiGormModels

import "back/LogInApi/config"

type Phone struct {
	ID     int    `gorm:"column:id"`
	UserId int    `gorm:"column:userid"`
	Phone  string `gorm:"column:phone"`
}

func (receiver Phone) TableName() string {
	return "phone"
}
func AddPhone(phone string, userId int) {
	config.DB.Create(&Phone{Phone: phone, UserId: userId})
}
func GetPhoneById(id int) Phone {
	var phone Phone
	config.DB.Where("id = ?", id).First(&phone)
	return phone
}
func GetPhoneByUserId(userId int) Phone {
	var phone Phone
	config.DB.Where("userid = ?", userId).First(&phone)
	return phone
}
func GetAllPhone() []Phone {
	var phone []Phone
	config.DB.Find(&phone)
	return phone
}
func UpdatePhone(phone Phone) {
	config.DB.Save(&phone)
}
func DeletePhone(phone Phone) {
	config.DB.Delete(&phone)
}
func GetUserIdFromPhone(phone string) int {
	var tmp Phone
	config.DB.Where("phone = ? ", phone).First(&tmp)
	if tmp.ID == 0 {
		return -1
	}
	return tmp.UserId
}
