package LogInApiGormModels

import (
	"back/LogInApi/config"

	"github.com/gin-gonic/gin"
)

type Password struct {
	ID     int    `gorm:"column:id"`
	UserId int    `gorm:"column:userid"`
	Pass   string `gorm:"column:password"`
}

func (receiver Password) TableName() string {
	return "password"
}
func AddPassword(pas string, userId int) {
	config.DB.Create(&Password{Pass: pas, UserId: userId})
}
func GetPasswordByUserId(userId int) []Password {
	var pas []Password
	config.DB.Where("userid = ?", userId).Last(&pas)
	return pas
}
func GetPasswordById(id int) Password {
	var pas Password
	config.DB.Where("id = ?", id).First(&pas)
	return pas
}
func AllPass(c *gin.Context) {
	var n []Password
	n = GetPasswordByUserId(2)
	c.IndentedJSON(200, n)
}
func GetAllPassword() []Password {
	var pas []Password
	config.DB.Find(&pas)
	return pas
}
func UpdatePassword(pas Password) {
	config.DB.Save(&pas)
}
func DeletePassword(pas Password) {
	config.DB.Delete(&pas)
}
func GetLastPasswordByUserId(userId int) Password {
	var pas Password
	//config.DB.Where("'userid' = ?", userId).Last(&pas)
	config.DB.Where("userid = ?", userId).Order("id DESC").First(&pas)
	return pas
}
