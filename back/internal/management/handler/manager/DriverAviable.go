package manager

import (
	authmodel "back/internal/auth/model"
	"back/internal/management/model"
	"back/internal/management/service"

	"github.com/gin-gonic/gin"
)

func ViewOrNotDriver(c *gin.Context) {
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
	// returns the requester's permission LEVEL, not their company id — see
	// the identical fix/comment in DeleteDriver.go.
	CompanyId := model.GetCompanyIdByLevelAndUserId(tmpDriver.UserId, tmpDriver.Level)
	ListOfDriversInCompany := model.GetAllDriversInCompany(CompanyId)
	for i := 0; i < len(ListOfDriversInCompany); i++ {

		if authmodel.GetFNameByUserID(ListOfDriversInCompany[i].UserId) == tmpDriver.DriverName {
			model.UpdateDriverAvailableValueByUserId(ListOfDriversInCompany[i].UserId)

			c.IndentedJSON(200, gin.H{"message": "good work"})
			return
		}
	}
	c.IndentedJSON(401, gin.H{})
	return
}
