package driver

import (
	"back/internal/management/handler"
	"back/internal/management/model"

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
	var tmp handler.UserIdAndLevel
	err := c.ShouldBindJSON(&tmp)
	if err != nil {
		c.IndentedJSON(401, gin.H{"message": "error in bind json"})
		return
	}
	// בדיקת הרשאות
	// NOTE: was service.ReturnPermissionForSingUp(tmp.UserId, tmp.Level),
	// which maps driver-level (5) permissions to 0 — the same value used
	// for "no permission" — so real drivers could never fetch their own
	// drives. That helper is meant for the admin/manager sign-up-elevation
	// check, not a general "is this a valid permission" gate; check the
	// permission row directly here instead.
	if model.GetPermissionByUserIdAndLevel(tmp.UserId, tmp.Level).Level == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}
	// קבלת מספר השיבוץ
	idCompany := model.GetCompanyIdByLevelAndUserId(tmp.UserId, tmp.Level)
	if idCompany == 0 {
		c.IndentedJSON(401, gin.H{"message": "you didnt have a company "})
		return
	}
	lastSortByCompanyId := model.GetNumOfLastSortByCompanyId(idCompany)
	if lastSortByCompanyId == 0 {
		c.IndentedJSON(401, gin.H{"message": "you didnt have a Sort "})
		return
	}
	userId := model.GetUserIdByCompanyIdAndLevel(tmp.Level, idCompany)
	if userId == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}
	driverId := model.GetDriverIdByUserIdAndCompanyId(userId, idCompany)
	tmpListDriveForDriver := model.GetAllDrivesByDriverID(driverId, lastSortByCompanyId)
	var tmpListToReturn []DriveToReturn
	var driveToReturn DriveToReturn
	for i := 0; i < len(tmpListDriveForDriver); i++ {
		driveToReturn = GetDriveToReturn(model.GetDriveById(tmpListDriveForDriver[i].DriveId))
		tmpListToReturn = append(tmpListToReturn, driveToReturn)
	}
	c.IndentedJSON(200, gin.H{"message": "work", "list of drives": tmpListToReturn})
	return
}
func GetDriveToReturn(drive model.Drive) DriveToReturn {
	return DriveToReturn{Name: drive.Name, HToStart: drive.HToStart, MToStart: drive.MToStart, GpsStart: drive.GpsStart,
		GpsEnd: drive.GpsEnd}
}
