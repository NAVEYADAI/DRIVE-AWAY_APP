package Manager

import (
	"back/MainController/MainControllerGormModels"
	"back/MainController/MainControllerLogic"
	"back/MainController/controller"

	"github.com/gin-gonic/gin"
)

// Use in DriverData Struct

//	type DriveData struct {
//		controller.UserIdAndLevel
//		DriveName string `json:"driveName"`
//	}
type DriveToReturn struct {
	NameOfDrive string `json:"nameOfDrive"`
	Aviable     bool   `json:"aviable"`
}

func NewDriveToReturn(NameOfDriver string, Aviable bool) DriveToReturn {
	return DriveToReturn{NameOfDrive: NameOfDriver, Aviable: Aviable}
}
func Get_Drives(c *gin.Context) {
	var tmpDrive controller.UserIdAndLevel
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

	idCompany := MainControllerGormModels.GetCompanyIdByLevelAndUserId(tmpDrive.UserId, tmpDrive.Level)
	listOfDrive := MainControllerGormModels.GetDriveAviablesByCompanyId(idCompany)
	var tmpReturnList []DriveToReturn
	var name string
	for i := 0; i < len(listOfDrive); i++ {
		name = MainControllerGormModels.GetDriveNameById(listOfDrive[i].IdDrive)
		tmpReturnList = append(tmpReturnList, NewDriveToReturn(name, listOfDrive[i].Available))
	}
	c.IndentedJSON(200, gin.H{"message": "Success", "List": tmpReturnList})

}
