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

type AddPosition struct {
	CompanyName string `json:"companyName"`

	handler.UserIdAndLevel

	handler.User
}

func Add_Position(c *gin.Context) {
	var tmpPosition AddPosition
	err := c.ShouldBindJSON(&tmpPosition)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in BindJSON"}) // 400
		return
	}
	// בדיקת הרשאות
	if service.ReturnPermissionForSingUp(tmpPosition.UserId, tmpPosition.Level) != 1 {
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
	CompanyId := model.GetCompanyByName(tmpPosition.CompanyName)
	tmpPermission := model.NewPermissions(UserId, CompanyId, 4)
	model.CreatePermissions(tmpPermission)

}

func copySingUpValuesPosition(tmpPosition AddPosition) (authhandler.SingUp, bool) {
	var singUp authhandler.SingUp
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
