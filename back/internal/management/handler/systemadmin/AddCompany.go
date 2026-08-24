package systemadmin

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

type AddCompany struct {
	CompanyName string `json:"companyName"`
	CompanyID   string `json:"companyID"`
	Address     string `json:"address"`

	handler.UserIdAndLevel

	handler.User
}

// localhost:2909/MainController/AddCompany
func Add_Company(c *gin.Context) {
	var tmpCompany AddCompany
	err := c.ShouldBindJSON(&tmpCompany)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"}) // 400
		return
	}

	// בדיקת הרשאות
	if service.ReturnPermissionForSingUp(tmpCompany.UserId, tmpCompany.Level) != 1 {
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

	// הוספת חברה
	good = goodeCompany(tmpCompany)
	if !good {
		c.IndentedJSON(400, gin.H{"message": "error in get value of company"}) // 400
		return
	}
	tmp := model.
		CreateCompanyWithReturnId(tmpCompany.CompanyName, tmpCompany.CompanyID, tmpCompany.Address)
	tmpPermission := model.NewPermissions(UserId, tmp.ID, 3)
	model.CreatePermissions(tmpPermission)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Company Created"})
	return
}

func copySingUpValuesCompany(tmpCompany AddCompany) (authhandler.SingUp, bool) {
	var singUp authhandler.SingUp
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
