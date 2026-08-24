package handler

import (
	authhandler "back/internal/auth/handler"
	authmodel "back/internal/auth/model"
	authservice "back/internal/auth/service"
	"back/internal/management/model"
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
	var logIn authhandler.LogIn
	err := c.ShouldBindJSON(&logIn)
	if err != nil {
		println(err)
		println("Error in BindJSON")
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "error in BindJSON"}) // 415
		authmodel.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 1)
		return
	}
	good, message := authservice.TryLogInIpAndMAcAddress(logIn.IP, logIn.MacAddress, time.Now())
	if !good {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message}) //409
		authmodel.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 2)
		return
	}
	UserId := authmodel.GetUserByUserName(logIn.UserName)
	if UserId == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"}) //404
		authmodel.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 3)
		return
	}
	pass := authmodel.GetLastPasswordByUserId(UserId)
	if pass.Pass != logIn.Password {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Password incorrect"}) //401
		authmodel.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 4)
		return
	}
	authmodel.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, true, time.Now(), -1)
	level := model.GetPermissionLevelByUserId(UserId)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "LogIn Success", "UserId": UserId, "level": level}) //200
}
