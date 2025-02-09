package Manager

import (
	"back/LogInApi/LogInApiGormModels"
	"back/MainController/MainControllerGormModels"
	"back/MainController/MainControllerLogic"

	"github.com/gin-gonic/gin"
)

// Use in DriverData Struct
//type DriverData struct {
//	controller.UserIdAndLevel
//	DriverName string `json:"driverName"`
//}

func Delete_Driver(c *gin.Context) {
	var tmpDriver DriverData
	err := c.BindJSON(&tmpDriver)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}

	// בדיקת הרשאות
	level := MainControllerLogic.ReturnPermissionForSingUp(tmpDriver.UserId, tmpDriver.Level)
	if level == -1 || level == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"})
		return
	}

	// קבלת המספר שורה של הנהג

	CompanyId := MainControllerGormModels.GetPermissionLevelByUserId(tmpDriver.UserId)
	ListOfDriversInCompany := MainControllerGormModels.GetAllDriversInCompany(CompanyId)
	for i := 0; i < len(ListOfDriversInCompany); i++ {

		if LogInApiGormModels.GetFNameByUserID(ListOfDriversInCompany[i].UserId) == tmpDriver.DriverName {
			MainControllerGormModels.DeleteDriverAvailableByDriverId(ListOfDriversInCompany[i].UserId)
			c.IndentedJSON(200, gin.H{"message": "Driver Is deleted"})
			return
		}
	}
}
