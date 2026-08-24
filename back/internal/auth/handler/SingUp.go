package handler

import (
	"back/internal/auth/model"
	"back/internal/auth/service"
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
	err := c.ShouldBindJSON(&singUp)
	if err != nil {
		println(err)
		println("Error in BindJSON")
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "error in BindJSON"})
		return
	}

	var goodPass bool
	good, message := service.SignUpNewUserName(singUp.UserName)
	if !good {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message})
	}
	goodPass, message = service.IsSecure(singUp.Password)
	if !goodPass {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message})
	}
	if !goodPass || !good {
		return
	}
	UserId := model.AddUser(singUp.UserName)
	service.AddPersonalInformation(singUp.FName, singUp.LName, singUp.IdentityCard, UserId)
	model.AddPersonalInformation(singUp.FName, singUp.LName, singUp.IdentityCard, UserId)
	model.AddPassword(singUp.Password, UserId)
	model.AddMail(singUp.Email, UserId)
	model.AddPhone(singUp.Phone, UserId)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User Created"})
}
func TmpSingUp(singUp SingUp) {
	var goodPass bool
	good, _ := service.SignUpNewUserName(singUp.UserName)
	if !good {
		print("user name is didnt good")
	}
	goodPass, _ = service.IsSecure(singUp.Password)
	if !goodPass {
		print("error in pass")
	}
	if !goodPass || !good {
		return
	}
	UserId := model.AddUser(singUp.UserName)
	model.AddPersonalInformation(singUp.FName, singUp.LName, singUp.IdentityCard, UserId)
	model.AddPassword(singUp.Password, UserId)
	model.AddMail(singUp.Email, UserId)
	model.AddPhone(singUp.Phone, UserId)
}
