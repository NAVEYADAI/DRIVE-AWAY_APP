package LogInApi

import (
	"back/LogInApi/LogInApiGormModels"
	"back/LogInApi/LogInApiLogic"
	"back/MainController/MainControllerGormModels"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LogInStruct struct {
	UserName   string `json:"UserName"`
	Password   string `json:"Password"`
	IP         string `json:"IP"`
	MacAddress string `json:"MacAddress"`
}

func Login(context *gin.Context) {
	var logIn LogInStruct
	err := context.BindJSON(&logIn)
	if err != nil {
		context.IndentedJSON(415, gin.H{"messege": "error in BindJSON"})
		LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 1)
		return
	}
	good, message := LogInApiLogic.TryLogInIpAndMAcAddress(logIn.IP, logIn.MacAddress, time.Now())
	if !good {
		context.IndentedJSON(http.StatusConflict, gin.H{"messege": message})
		LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 2)
		return
	}
	UserId := LogInApiGormModels.GetUserByUserName(logIn.UserName)
	if UserId == -1 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"messege": "User not found"})
		LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 3)
		return
	}
	pass := LogInApiGormModels.GetLastPasswordByUserId(UserId)
	if pass.Pass != logIn.Password {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"messege": "Password incorrect"})
		LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 4)
		return
	}
	LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, true, time.Now(), -1)
	perm := MainControllerGormModels.GetPermissionsByUserId(UserId)
	if perm.GetId() != 0 {
		context.IndentedJSON(http.StatusOK, gin.H{"messege": "LogIn Success", "UserId": UserId, "level": perm.GetLevel()})
		return
	}
	context.IndentedJSON(http.StatusUnauthorized, gin.H{"messege": "User has no permissions"})
}
