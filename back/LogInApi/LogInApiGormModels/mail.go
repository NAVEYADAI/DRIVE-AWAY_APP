package LogInApiGormModels

import "back/LogInApi/config"

type Mail struct {
	ID     int    `gorm:"column:id"`
	UserId int    `gorm:"column:userid"`
	Mail   string `gorm:"column:mail"`
}

func (receiver Mail) TableName() string {
	return "mail"
}

func AddMail(mail string, userId int) {
	config.DB.Create(&Mail{Mail: mail, UserId: userId})
}

func GetMailByUserId(userId int) Mail {
	var mail Mail
	config.DB.Where("userid = ?", userId).First(&mail)
	return mail
}
func GetAllMail() []Mail {
	var mail []Mail
	config.DB.Find(&mail)
	return mail
}
func UpdateMail(mail Mail) {
	config.DB.Save(&mail)
}
func DeleteMail(mail Mail) {
	config.DB.Delete(&mail)
}

func GetUserIdFromMail(mail string) int {
	var tmp Mail
	config.DB.Where("mail = ?", mail).First(&tmp)
	if tmp.ID == 0 {
		return -1
	}
	return tmp.UserId
}
func GetMailById(id int) Mail {
	var mail Mail
	config.DB.Where("id = ?", id).First(&mail)
	return mail
}
