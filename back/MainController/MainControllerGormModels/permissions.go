package MainControllerGormModels

import (
	"back/MainController/config"
)

type Permissions struct {
	ID        int `gorm:"column:id"`
	UserId    int `gorm:"column:userid"`
	ConpanyId int `gorm:"column:companyid"`
	Level     int `gorm:"column:level"`
}

func (receiver Permissions) TableName() string {
	return "permissions"
}
func NewPermissions(UserId, CompanyId, level int) Permissions {
	return Permissions{UserId: UserId, ConpanyId: CompanyId, Level: level}
}
func (p Permissions) GetId() int {
	return p.ID
}
func (p Permissions) GetUserId() int {
	return p.UserId
}
func (p Permissions) GetCompanyId() int {
	return p.ConpanyId
}
func (p Permissions) GetLevel() int {
	return p.Level
}
func (p *Permissions) SetUserId(userId int) {
	p.UserId = userId
}
func (p *Permissions) SetCompanyId(companyId int) {
	p.ConpanyId = companyId
}
func (p *Permissions) SetLevel(level int) {
	p.Level = level
}

// Gorm functions
func CreatePermissions(p Permissions) {
	config.DB.Create(&p)
}
func GetAllPermissions() []Permissions {
	var permissions []Permissions
	config.DB.Find(&permissions)
	return permissions
}
func GetPermissionLevelByUserId(userId int) int {
	var permissions Permissions
	config.DB.Where("userid = ?", userId).Find(&permissions)
	return permissions.Level
}
func GetPermissionsById(id int) Permissions {
	var permission Permissions
	config.DB.Where("ID = ?", id).Find(&permission)
	return permission
}
func GetPermissionsByUserId(userId int) Permissions {
	var permissions Permissions
	err := config.DB.Where("userid = ?", userId).Find(&permissions)
	if err.Error != nil {
		print(err.Error)
	}
	return permissions
}
func GetPermissionsByAllId(userId, companyId, level int) int {
	var permissions Permissions
	config.DB.Where("userid = ? AND companyid = ? AND level = ?", userId, companyId, level).Find(&permissions)
	return permissions.ID
}

func UpdatePermissions(p Permissions) {
	config.DB.Save(&p)
}
func DeletePermissions(p Permissions) {
	config.DB.Delete(&p)
}
func GetCompanyIdByLevelAndUserId(userId, level int) int {
	var permissions Permissions
	config.DB.Where("userid = ? AND level = ?", userId, level).Find(&permissions)
	return permissions.ConpanyId
}
func GetPermissionByUserIdAndLevel(userId, level int) Permissions {
	var permissions Permissions
	// ביצוע השאילתה עם ניפוי שגיאות
	//query := config.DB.Where("userid = ? AND level = ?", userId, level)
	//result := query.Find(&permissions)
	//

	//result := config.DB.Where("userid = ? AND level = ?", userId, level).Find(&permissions)

	config.DB.Where("userid = ? AND level = ?", userId, level).Find(&permissions)

	// בדיקת תוצאה
	//fmt.Println("Rows affected:", result.RowsAffected)
	//if result.Error != nil {
	//	fmt.Println("Error:", result.Error)
	//	log.Fatal(result.Error)
	//}
	//
	//fmt.Println("Permissions found:", permissions)

	return permissions
}
func GetUserIdByCompanyIdAndLevel(level, companyId int) int {
	var tmp Permissions
	config.DB.Where("level = ? AND companyid = ?", level, companyId).Find(&tmp)
	return tmp.UserId
}
