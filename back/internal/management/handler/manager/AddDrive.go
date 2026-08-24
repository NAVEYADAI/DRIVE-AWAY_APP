package manager

import (
	"back/internal/management/model"
	"back/internal/management/service"
	"back/internal/management/handler"

	"github.com/gin-gonic/gin"
)

type AddDrive struct {
	NameOfDriver string `json:"nameOfDriver"`
	TimeToStartH int    `json:"timeToStartH"`
	TimeToStartM int    `json:"timeToStartM"`
	GpsStart     string `json:"gpsStart"`
	GpsEnd       string `json:"gpsEnd"`

	handler.UserIdAndLevel
}

func Add_Drive(c *gin.Context) {
	var tmpDrive AddDrive
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

	// קבלת החברה של המנהל חברה
	CompanyId := model.GetCompanyIdByLevelAndUserId(tmpDrive.UserId, tmpDrive.Level)

	// הוספת נסיעה
	drive := model.NewDrive(tmpDrive.NameOfDriver, tmpDrive.TimeToStartH, tmpDrive.TimeToStartM, CompanyId, tmpDrive.GpsStart, tmpDrive.GpsEnd)
	DriveId := model.CreateDriveReturnId(drive)

	// הוספת זמינות של הנהג לטבלה
	driveAvailability := model.NewDriveAvailable(CompanyId, DriveId, true)
	model.CreateDriveAvailable(driveAvailability)
}
