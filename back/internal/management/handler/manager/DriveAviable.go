package manager

import (
	"back/internal/management/model"
	"back/internal/management/service"

	"github.com/gin-gonic/gin"
)

func ViewOrNotDrive(c *gin.Context) {
	var tmpDrive DriveData
	err := c.ShouldBindJSON(&tmpDrive)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"}) // 400
		return
	}

	// בדיקת הרשאות
	level := service.ReturnPermissionForSingUp(tmpDrive.UserId, tmpDrive.Level)
	if level == -1 || level == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}

	//קבלת המספר שורה של הנסיעה
	idDrive := model.GetDriveByName(tmpDrive.DriveName)
	idCompany := model.GetCompanyIdByLevelAndUserId(tmpDrive.UserId, tmpDrive.Level)
	// Same company-ownership check as Delete_Drive: don't let a manager from
	// one company toggle another company's drive availability.
	drive := model.GetDriveById(idDrive)
	if drive.CompanyId == idCompany && 0 != model.GetPermissionsByAllId(tmpDrive.UserId, idCompany, tmpDrive.Level) {
		model.UpdateDriveAvailableValueByDriveId(idDrive)
		c.IndentedJSON(200, gin.H{"message": "good work"})
		return
	}
	c.IndentedJSON(401, gin.H{})
	return
}
