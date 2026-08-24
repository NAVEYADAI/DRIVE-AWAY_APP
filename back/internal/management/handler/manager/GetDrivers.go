package manager

import (
	authmodel "back/internal/auth/model"
	"back/internal/management/handler"
	"back/internal/management/model"
	"back/internal/management/service"

	"github.com/gin-gonic/gin"
)

type DriverToReturn struct {
	NameOfDriver string `json:"nameOfDriver"`
	Aviable      bool   `json:"aviable"`
}

func NewDriverToReturn(NameOfDriver string, Aviable bool) DriverToReturn {
	return DriverToReturn{NameOfDriver: NameOfDriver, Aviable: Aviable}
}
func Get_Drivers(c *gin.Context) {
	var tmpDriver handler.UserIdAndLevel
	err := c.ShouldBindJSON(&tmpDriver)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"}) // 400
		return
	}

	// בדיקת הרשאות
	level := service.ReturnPermissionForSingUp(tmpDriver.UserId, tmpDriver.Level)
	if level == -1 || level == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}

	idCompany := model.GetCompanyIdByLevelAndUserId(tmpDriver.UserId, tmpDriver.Level)
	listOfDrivers := model.GetAllDriversAviableInCompany(idCompany)
	var tmpReturnList []DriverToReturn
	var name string
	for i := 0; i < len(listOfDrivers); i++ {
		name = authmodel.GetFNameByUserID(listOfDrivers[i].IdDriver)
		tmpReturnList = append(tmpReturnList, NewDriverToReturn(name, listOfDrivers[i].Available))
	}
	c.IndentedJSON(200, gin.H{"message": "Success", "List": tmpReturnList})
	//c.Data(http.StatusOK, "listOfDriver", tmpReturnList)
}
