package model

import (
	"back/internal/config"
	"strconv"
)

type PersonalInformation struct {
	ID           int    `gorm:"column:id"`
	FName        string `gorm:"column:fname"`
	LName        string `gorm:"column:lname"`
	IdentityCard string `gorm:"column:identitycard"`
	UserId       int    `gorm:"column:userid"`
}

func (receiver PersonalInformation) TableName() string {
	return "personalinformation"
}
func AddPersonalInformation(fname string, lname string, identitycard string, userid int) {
	config.DB.Create(&PersonalInformation{FName: fname, LName: lname, IdentityCard: identitycard, UserId: userid})
}
func GetPersonalInformationById(id int) PersonalInformation {
	var personalInformation PersonalInformation
	config.DB.Where("id = ?", id).First(&personalInformation)
	return personalInformation
}
func GetAllPersonalInformation() []PersonalInformation {
	var personalInformation []PersonalInformation
	config.DB.Find(&personalInformation)
	return personalInformation
}
func GetFNameByUserID(userId int) string {
	var tmpPersonalInformation PersonalInformation
	config.DB.Where("userid", userId).Find(&tmpPersonalInformation)
	return tmpPersonalInformation.FName
}

func UpdatePersonalInformation(personalInformation PersonalInformation) {
	config.DB.Save(&personalInformation)
}
func DeletePersonalInformation(personalInformation PersonalInformation) {
	config.DB.Delete(&personalInformation)
}
func GetUserIdFromID(ID int) int {
	var tmp PersonalInformation
	// IdentityCard is stored as text; comparing it against a bare int
	// parameter fails to encode at the driver level (identitycard is OID
	// text, not numeric), so the value must be stringified first.
	config.DB.Where("identitycard = ? ", strconv.Itoa(ID)).First(&tmp)
	if tmp.ID == 0 {
		return -1
	}
	return tmp.UserId
}
func GetPersonalInformationByIDCard(IDCard string) PersonalInformation {
	var tmp PersonalInformation
	config.DB.Where("identitycard = ? ", IDCard).First(&tmp)
	return tmp
}
