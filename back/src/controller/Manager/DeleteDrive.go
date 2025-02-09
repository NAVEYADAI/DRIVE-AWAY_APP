package Manager

import (
	"back/MainController/MainControllerGormModels"
	"back/MainController/MainControllerLogic"

	"github.com/gin-gonic/gin"
)

// Use in DriveData Struct

func Delete_Drive(c *gin.Context) {
	var tmpDrive DriveData
	err := c.BindJSON(&tmpDrive)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"}) // 400
		return
	}

	// בדיקת הרשאות
	level := MainControllerLogic.ReturnPermissionForSingUp(tmpDrive.UserId, tmpDrive.Level)
	if level == -1 || level == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}

	//קבלת המספר שורה של הנסיעה
	idDrive := MainControllerGormModels.GetDriveByName(tmpDrive.DriveName)
	idCompany := MainControllerGormModels.GetCompanyIdByLevelAndUserId(tmpDrive.UserId, tmpDrive.Level)
	if 0 != MainControllerGormModels.GetPermissionsByAllId(tmpDrive.UserId, idCompany, tmpDrive.Level) {
		MainControllerGormModels.DeleteDriveAvailableByDriveId(idDrive)
		c.IndentedJSON(200, gin.H{"message": "Drive deleted"}) // 200
		return
	}
	c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
	return

}
