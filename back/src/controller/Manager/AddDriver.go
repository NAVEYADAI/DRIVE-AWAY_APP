package Manager

import (
	"back/LogInApi/LogInApiController"
	"back/LogInApi/LogInApiGormModels"
	"back/LogInApi/LogInApiLogic"
	"back/MainController/MainControllerGormModels"
	"back/MainController/MainControllerLogic"
	"back/MainController/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddDriver struct {
	controller.UserIdAndLevel

	controller.User

	Address    string `json:"address"`
	IsEmployee bool   `json:"isEmployee"`
}

func Add_Driver(c *gin.Context) {
	var tmpDriver AddDriver
	err := c.BindJSON(&tmpDriver)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"}) // 400
		return
	}

	// בדיקת הרשאות
	level := MainControllerLogic.ReturnPermissionForSingUp(tmpDriver.UserId, tmpDriver.Level)
	if level == -1 || level == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}

	// הוספת משתמש
	singUp, work := copySingUpValuesAddDriver(tmpDriver)

	if !work {
		c.IndentedJSON(400, gin.H{"message": "error in get value"}) // 400
		return
	}
	var goodPass bool
	good, message := LogInApiLogic.SignUpNewUserName(singUp.UserName)
	if !good {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message})
	}
	goodPass, message = LogInApiLogic.IsSecure(singUp.Password)
	if !goodPass {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message})
	}
	if !goodPass || !good {
		return
	}
	UserId := LogInApiGormModels.AddUser(singUp.UserName)
	LogInApiLogic.AddPersonalInformation(singUp.FName, singUp.LName, singUp.IdentityCard, UserId)
	LogInApiGormModels.AddPersonalInformation(singUp.FName, singUp.LName, singUp.IdentityCard, UserId)
	LogInApiGormModels.AddPassword(singUp.Password, UserId)
	LogInApiGormModels.AddMail(singUp.Email, UserId)
	LogInApiGormModels.AddPhone(singUp.Phone, UserId)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User Created"})

	// הוספת הרשאה
	CompanyId := MainControllerGormModels.GetCompanyIdByLevelAndUserId(tmpDriver.UserId, tmpDriver.Level)
	tmpPermission := MainControllerGormModels.NewPermissions(UserId, CompanyId, 5)
	MainControllerGormModels.CreatePermissions(tmpPermission)

	MainControllerGormModels.CreateDriver(MainControllerGormModels.NewDriver(UserId, CompanyId, tmpDriver.Address, tmpDriver.IsEmployee))

	MainControllerGormModels.CreateDriverAvailable(MainControllerGormModels.NewDriverAvailable(CompanyId, UserId, true))

}

func copySingUpValuesAddDriver(tmp AddDriver) (LogInApiController.SingUp, bool) {
	var singUp LogInApiController.SingUp
	if tmp.UserName == "" ||
		tmp.FName == "" ||
		tmp.LName == "" ||
		tmp.IdentityCard == "" ||
		tmp.Password == "" ||
		tmp.Email == "" ||
		tmp.Phone == "" {
		return singUp, false
	}
	singUp.UserName = tmp.UserName
	singUp.FName = tmp.FName
	singUp.LName = tmp.LName
	singUp.IdentityCard = tmp.IdentityCard
	singUp.Password = tmp.Password
	singUp.Email = tmp.Email
	singUp.Phone = tmp.Phone
	return singUp, true
}
