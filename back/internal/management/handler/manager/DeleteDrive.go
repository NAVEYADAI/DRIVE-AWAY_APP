package manager

import (
	"back/internal/management/model"
	"back/internal/management/service"

	"github.com/gin-gonic/gin"
)

// Use in DriveData Struct

func Delete_Drive(c *gin.Context) {
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
	// The drive being deleted must actually belong to the requester's own
	// company — otherwise any manager could delete another company's drive
	// just by knowing/guessing its name (GetDriveByName isn't company-scoped).
	drive := model.GetDriveById(idDrive)
	if drive.CompanyId == idCompany && 0 != model.GetPermissionsByAllId(tmpDrive.UserId, idCompany, tmpDrive.Level) {
		model.DeleteDriveAvailableByDriveId(idDrive)
		c.IndentedJSON(200, gin.H{"message": "Drive deleted"}) // 200
		return
	}
	c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
	return

}
