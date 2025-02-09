package LogInApiController

import (
	"back/LogInApi/LogInApiGormModels"
	"back/LogInApi/LogInApiLogic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SingUp struct {
	UserName     string `json:"UserName"`
	FName        string `json:"FName"`
	LName        string `json:"LName"`
	IdentityCard string `json:"IdentityCard"`
	Password     string `json:"Password"`
	Email        string `json:"Email"`
	Phone        string `json:"Phone"`
}

func CreateUser(c *gin.Context) {
	var singUp SingUp
	err := c.BindJSON(&singUp)
	if err != nil {
		println(err)
		println("Error in BindJSON")
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "error in BindJSON"})
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
}
func TmpSingUp(singUp SingUp) {
	var goodPass bool
	good, _ := LogInApiLogic.SignUpNewUserName(singUp.UserName)
	if !good {
		print("user name is didnt good")
	}
	goodPass, _ = LogInApiLogic.IsSecure(singUp.Password)
	if !goodPass {
		print("error in pass")
	}
	if !goodPass || !good {
		return
	}
	UserId := LogInApiGormModels.AddUser(singUp.UserName)
	LogInApiGormModels.AddPersonalInformation(singUp.FName, singUp.LName, singUp.IdentityCard, UserId)
	LogInApiGormModels.AddPassword(singUp.Password, UserId)
	LogInApiGormModels.AddMail(singUp.Email, UserId)
	LogInApiGormModels.AddPhone(singUp.Phone, UserId)
}
