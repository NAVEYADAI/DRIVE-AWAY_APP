package LogInApiGormModels

import (
	"back/LogInApi/config"
	"log"
)

type Users struct {
	ID       int    `gorm:"column:id"`
	UserName string `gorm:"column:userName"`
}

func (receiver Users) TableName() string {
	return "users"
}

func AddUser(userName string) int {
	user := &Users{UserName: userName}
	config.DB.Create(user)
	return user.ID
}
func GetUserById(id int) Users {
	var user Users
	config.DB.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return Users{ID: -1}
	}
	return user
}

func GetAllUser() []Users {
	var users []Users
	if err := config.DB.Find(&users).Error; err != nil {
		log.Println("Failed to retrieve users:", err)
		return nil
	}
	return users
}
func UpdateUser(user string, id int) {
	config.DB.Save(&Users{UserName: user, ID: id})
}
func DeleteUser(user Users) {
	config.DB.Delete(&user)
}
func GetUserByUserName(userName string) int {
	var user Users
	config.DB.Where("userName", userName).Find(&user)
	return user.ID
}
