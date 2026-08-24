package manager

import (
	authmodel "back/internal/auth/model"
	"back/internal/management/model"
	"back/internal/management/service"

	"github.com/gin-gonic/gin"
)

// Use in DriverData Struct
//type DriverData struct {
//	controller.UserIdAndLevel
//	DriverName string `json:"driverName"`
//}

func Delete_Driver(c *gin.Context) {
	var tmpDriver DriverData
	err := c.ShouldBindJSON(&tmpDriver)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}

	// בדיקת הרשאות
	level := service.ReturnPermissionForSingUp(tmpDriver.UserId, tmpDriver.Level)
	if level == -1 || level == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"})
		return
	}

	// קבלת המספר שורה של הנהג

	// NOTE: was model.GetPermissionLevelByUserId(tmpDriver.UserId), which
	// returns the requester's permission LEVEL, not their company id — that
	// value was then used as a company id, so this almost never found the
	// right company's drivers. See DeleteDriver_test.go.
	CompanyId := model.GetCompanyIdByLevelAndUserId(tmpDriver.UserId, tmpDriver.Level)
	ListOfDriversInCompany := model.GetAllDriversInCompany(CompanyId)
	for i := 0; i < len(ListOfDriversInCompany); i++ {

		if authmodel.GetFNameByUserID(ListOfDriversInCompany[i].UserId) == tmpDriver.DriverName {
			model.DeleteDriverAvailableByDriverId(ListOfDriversInCompany[i].UserId)
			c.IndentedJSON(200, gin.H{"message": "Driver Is deleted"})
			return
		}
	}
}
