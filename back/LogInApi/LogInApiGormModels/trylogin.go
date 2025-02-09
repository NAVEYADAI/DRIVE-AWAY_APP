package LogInApiGormModels

import (
	"back/LogInApi/config"
	"time"
)

/*
Cause
-1 - secsseful
1 - user not in context
2 - Too many attempts
3 - Username not found
4 - different password
*/
type TryLogIn struct {
	ID         int       `gorm:"column:id"`
	UserName   string    `gorm:"column:userName"`
	Password   string    `gorm:"column:password"`
	IP         string    `gorm:"column:ip"`
	MacAddress string    `gorm:"column:macAddress"`
	Cause      int       `gorm:"column:Cause"`
	Connect    bool      `gorm:"column:connect"`
	Date       time.Time `gorm:"column:date"`
}

func (receiver TryLogIn) TableName() string {
	return "trylogin"
}

func AddTryLogIn(userName string, password string, ip string, macAddress string, connect bool, date time.Time, Cause int) {
	config.DB.Create(&TryLogIn{UserName: userName, Password: password, IP: ip, MacAddress: macAddress, Connect: connect, Date: date, Cause: Cause})
}
func GetTryLogInById(id int) TryLogIn {
	var tryLogIn TryLogIn
	config.DB.Where("id = ?", id).First(&tryLogIn)
	return tryLogIn
}
func GetAllTryLogIn() []TryLogIn {
	var tryLogIn []TryLogIn
	config.DB.Find(&tryLogIn)
	return tryLogIn
}
func GetTryLoginById(id int) TryLogIn {
	var tryLogIn TryLogIn
	config.DB.Where("id = ?", id).First(&tryLogIn)
	return tryLogIn
}
func UpdateTryLogIn(tryLogIn TryLogIn) {
	config.DB.Save(&tryLogIn)
}
func DeleteTryLogIn(tryLogIn TryLogIn) {
	config.DB.Delete(&tryLogIn)
}
func GetTryLogInByIpAndMacAddress(ip string, macAddress string, t time.Time) []TryLogIn {
	var tryLogIn []TryLogIn
	//newTime := t.Add(-10 * time.Minute)
	config.DB.Where("ip = ? AND 'macAddress' = ? ", ip, macAddress).Find(&tryLogIn)
	//config.DB.Where("IP = ? AND macAddress = ? AND data < ?", ip, macAddress, newTime).Find(&tryLogIn)
	//
	return tryLogIn
}
