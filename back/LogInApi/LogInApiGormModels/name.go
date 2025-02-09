package LogInApiGormModels

import "back/LogInApi/config"

type Name struct {
	ID     int    `gorm:"column:id"`
	IDUser int    `gorm:"column:idUser"`
	Name   string `gorm:"column:name"`
}

func (receiver Name) TableName() string {
	return "name"
}

func AddName(name string, idUser int) {
	config.DB.Create(&Name{Name: name, IDUser: idUser})
}
func GetNameById(id int) Name {
	var name Name
	config.DB.Where("id = ?", id).First(&name)
	return name
}
func GetAllName() []Name {
	var name []Name
	config.DB.Find(&name)
	return name
}
func UpdateName(name Name) {
	config.DB.Save(&name)
}
func DeleteName(name Name) {
	config.DB.Delete(&name)
}
func GetUserListById(id int) []Name {
	var user []Name
	config.DB.Where("idUser = ?", id).Find(&user)
	return user
}
