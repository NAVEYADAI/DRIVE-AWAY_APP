package controller

import (
	"back/LogInApi/LogInApiController"
	"back/LogInApi/LogInApiGormModels"
	"back/LogInApi/LogInApiLogic"
	"back/MainController/MainControllerGormModels"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//	type LogIn struct {
//		UserName   string `json:"UserName"`
//		Password   string `json:"Password"`
//		IP         string `json:"IP"`
//		MacAddress string `json:"MacAddress"`
//	}
//

// localhost:2909/MainController/LogIn
func LogIn(c *gin.Context) {
	var logIn LogInApiController.LogIn
	err := c.BindJSON(&logIn)
	if err != nil {
		println(err)
		println("Error in BindJSON")
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "error in BindJSON"}) // 415
		LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 1)
		return
	}
	good, message := LogInApiLogic.TryLogInIpAndMAcAddress(logIn.IP, logIn.MacAddress, time.Now())
	if !good {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message}) //409
		LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 2)
		return
	}
	UserId := LogInApiGormModels.GetUserByUserName(logIn.UserName)
	if UserId == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"}) //404
		LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 3)
		return
	}
	pass := LogInApiGormModels.GetLastPasswordByUserId(UserId)
	if pass.Pass != logIn.Password {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Password incorrect"}) //401
		LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 4)
		return
	}
	LogInApiGormModels.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, true, time.Now(), -1)
	level := MainControllerGormModels.GetPermissionLevelByUserId(UserId)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "LogIn Success", "UserId": UserId, "level": level}) //200
}
