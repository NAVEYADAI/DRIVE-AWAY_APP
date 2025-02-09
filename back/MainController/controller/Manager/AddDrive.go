package Manager

import (
	"back/MainController/MainControllerGormModels"
	"back/MainController/MainControllerLogic"
	"back/MainController/controller"

	"github.com/gin-gonic/gin"
)

type AddDrive struct {
	NameOfDriver string `json:"nameOfDriver"`
	TimeToStartH int    `json:"timeToStartH"`
	TimeToStartM int    `json:"timeToStartM"`
	GpsStart     string `json:"gpsStart"`
	GpsEnd       string `json:"gpsEnd"`

	controller.UserIdAndLevel
}

func Add_Drive(c *gin.Context) {
	var tmpDrive AddDrive
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

	// קבלת החברה של המנהל חברה
	CompanyId := MainControllerGormModels.GetCompanyIdByLevelAndUserId(tmpDrive.UserId, tmpDrive.Level)

	// הוספת נסיעה
	drive := MainControllerGormModels.NewDrive(tmpDrive.NameOfDriver, tmpDrive.TimeToStartH, tmpDrive.TimeToStartM, CompanyId, tmpDrive.GpsStart, tmpDrive.GpsEnd)
	DriveId := MainControllerGormModels.CreateDriveReturnId(drive)

	// הוספת זמינות של הנהג לטבלה
	driveAvailability := MainControllerGormModels.NewDriveAvailable(CompanyId, DriveId, true)
	MainControllerGormModels.CreateDriveAvailable(driveAvailability)
}
