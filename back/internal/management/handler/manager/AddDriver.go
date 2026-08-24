package manager

import (
	authhandler "back/internal/auth/handler"
	authmodel "back/internal/auth/model"
	authservice "back/internal/auth/service"
	"back/internal/management/handler"
	"back/internal/management/model"
	"back/internal/management/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddDriver struct {
	handler.UserIdAndLevel

	handler.User

	Address    string `json:"address"`
	IsEmployee bool   `json:"isEmployee"`
}

func Add_Driver(c *gin.Context) {
	var tmpDriver AddDriver
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

	// הוספת משתמש
	singUp, work := copySingUpValuesAddDriver(tmpDriver)

	if !work {
		c.IndentedJSON(400, gin.H{"message": "error in get value"}) // 400
		return
	}
	var goodPass bool
	good, message := authservice.SignUpNewUserName(singUp.UserName)
	if !good {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message})
	}
	goodPass, message = authservice.IsSecure(singUp.Password)
	if !goodPass {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message})
	}
	if !goodPass || !good {
		return
	}
	UserId := authmodel.AddUser(singUp.UserName)
	authservice.AddPersonalInformation(singUp.FName, singUp.LName, singUp.IdentityCard, UserId)
	authmodel.AddPersonalInformation(singUp.FName, singUp.LName, singUp.IdentityCard, UserId)
	authmodel.AddPassword(singUp.Password, UserId)
	authmodel.AddMail(singUp.Email, UserId)
	authmodel.AddPhone(singUp.Phone, UserId)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User Created"})

	// הוספת הרשאה
	CompanyId := model.GetCompanyIdByLevelAndUserId(tmpDriver.UserId, tmpDriver.Level)
	tmpPermission := model.NewPermissions(UserId, CompanyId, 5)
	model.CreatePermissions(tmpPermission)

	model.CreateDriver(model.NewDriver(UserId, CompanyId, tmpDriver.Address, tmpDriver.IsEmployee))

	model.CreateDriverAvailable(model.NewDriverAvailable(CompanyId, UserId, true))

}

func copySingUpValuesAddDriver(tmp AddDriver) (authhandler.SingUp, bool) {
	var singUp authhandler.SingUp
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
