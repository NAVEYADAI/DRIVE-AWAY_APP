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

type AddCompany struct {
	CompanyName string `json:"companyName"`
	CompanyID   string `json:"companyID"`
	Address     string `json:"address"`

	controller.UserIdAndLevel

	controller.User
}

// localhost:2909/MainController/AddCompany
func Add_Company(c *gin.Context) {
	var tmpCompany AddCompany
	err := c.BindJSON(&tmpCompany)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"}) // 400
		return
	}

	// בדיקת הרשאות
	if MainControllerLogic.ReturnPermissionForSingUp(tmpCompany.UserId, tmpCompany.Level) != 1 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}

	// הוספת משתמש
	singUp, work := copySingUpValuesCompany(tmpCompany)

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

	// הוספת חברה
	good = goodeCompany(tmpCompany)
	if !good {
		c.IndentedJSON(400, gin.H{"message": "error in get value of company"}) // 400
		return
	}
	tmp := MainControllerGormModels.
		CreateCompanyWithReturnId(tmpCompany.CompanyName, tmpCompany.CompanyID, tmpCompany.Address)
	tmpPermission := MainControllerGormModels.NewPermissions(UserId, tmp.ID, 3)
	MainControllerGormModels.CreatePermissions(tmpPermission)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Company Created"})
	return
}

func copySingUpValuesCompany(tmpCompany AddCompany) (LogInApiController.SingUp, bool) {
	var singUp LogInApiController.SingUp
	if tmpCompany.UserName == "" ||
		tmpCompany.FName == "" ||
		tmpCompany.LName == "" ||
		tmpCompany.IdentityCard == "" ||
		tmpCompany.Password == "" ||
		tmpCompany.Email == "" ||
		tmpCompany.Phone == "" {
		return singUp, false
	}
	singUp.UserName = tmpCompany.UserName
	singUp.FName = tmpCompany.FName
	singUp.LName = tmpCompany.LName
	singUp.IdentityCard = tmpCompany.IdentityCard
	singUp.Password = tmpCompany.Password
	singUp.Email = tmpCompany.Email
	singUp.Phone = tmpCompany.Phone
	return singUp, true
}
func goodeCompany(tmp AddCompany) bool {
	if tmp.CompanyName == "" || tmp.CompanyID == "" || tmp.Address == "" {
		return false
	}
	return true

}
