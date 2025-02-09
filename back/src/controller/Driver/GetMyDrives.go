package Driver

import (
	"back/MainController/MainControllerGormModels"
	"back/MainController/MainControllerLogic"
	"back/MainController/controller"

	"github.com/gin-gonic/gin"
)

type DriveToReturn struct {
	Name     string `json:"Name"`
	HToStart int    `json:"HToStart"`
	MToStart int    `json:"MToStart"`
	GpsStart string `json:"GpsStart"`
	GpsEnd   string `json:"GpsEnd"`
}

func GetMyDrives(c *gin.Context) {
	var tmp controller.UserIdAndLevel
	err := c.BindJSON(&tmp)
	if err != nil {
		c.IndentedJSON(401, gin.H{"message": "error in bind json"})
		return
	}
	// בדיקת הרשאות
	level := MainControllerLogic.ReturnPermissionForSingUp(tmp.UserId, tmp.Level)
	if level == -1 || level == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}
	// קבלת מספר השיבוץ
	idCompany := MainControllerGormModels.GetCompanyIdByLevelAndUserId(tmp.UserId, tmp.Level)
	if idCompany == 0 {
		c.IndentedJSON(401, gin.H{"message": "you didnt have a company "})
		return
	}
	lastSortByCompanyId := MainControllerGormModels.GetNumOfLastSortByCompanyId(idCompany)
	if lastSortByCompanyId == 0 {
		c.IndentedJSON(401, gin.H{"message": "you didnt have a Sort "})
		return
	}
	userId := MainControllerGormModels.GetUserIdByCompanyIdAndLevel(tmp.Level, idCompany)
	if userId == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}
	driverId := MainControllerGormModels.GetDriverIdByUserIdAndCompanyId(userId, idCompany)
	tmpListDriveForDriver := MainControllerGormModels.GetAllDrivesByDriverID(driverId, lastSortByCompanyId)
	var tmpListToReturn []DriveToReturn
	var driveToReturn DriveToReturn
	for i := 0; i < len(tmpListDriveForDriver); i++ {
		driveToReturn = GetDriveToReturn(MainControllerGormModels.GetDriveById(tmpListDriveForDriver[i].DriveId))
		tmpListToReturn = append(tmpListToReturn, driveToReturn)
	}
	c.IndentedJSON(200, gin.H{"message": "work", "list of drives": tmpListToReturn})
	return
}
func GetDriveToReturn(drive MainControllerGormModels.Drive) DriveToReturn {
	return DriveToReturn{Name: drive.Name, HToStart: drive.HToStart, MToStart: drive.MToStart, GpsStart: drive.GpsStart,
		GpsEnd: drive.GpsEnd}
}
