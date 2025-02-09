package SystemAdministrator

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

type AddPosition struct {
	CompanyName string `json:"companyName"`

	controller.UserIdAndLevel

	controller.User
}

func Add_Position(c *gin.Context) {
	var tmpPosition AddPosition
	err := c.BindJSON(&tmpPosition)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"}) // 400
		return
	}
	// בדיקת הרשאות
	if MainControllerLogic.ReturnPermissionForSingUp(tmpPosition.UserId, tmpPosition.Level) != 1 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}

	// הוספת משתמש
	singUp, work := copySingUpValuesPosition(tmpPosition)

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
	CompanyId := MainControllerGormModels.GetCompanyByName(tmpPosition.CompanyName)
	tmpPermission := MainControllerGormModels.NewPermissions(UserId, CompanyId, 4)
	MainControllerGormModels.CreatePermissions(tmpPermission)

}

func copySingUpValuesPosition(tmpPosition AddPosition) (LogInApiController.SingUp, bool) {
	var singUp LogInApiController.SingUp
	if tmpPosition.UserName == "" ||
		tmpPosition.FName == "" ||
		tmpPosition.LName == "" ||
		tmpPosition.IdentityCard == "" ||
		tmpPosition.Password == "" ||
		tmpPosition.Email == "" ||
		tmpPosition.Phone == "" {
		return singUp, false
	}
	singUp.UserName = tmpPosition.UserName
	singUp.FName = tmpPosition.FName
	singUp.LName = tmpPosition.LName
	singUp.IdentityCard = tmpPosition.IdentityCard
	singUp.Password = tmpPosition.Password
	singUp.Email = tmpPosition.Email
	singUp.Phone = tmpPosition.Phone
	return singUp, true
}
