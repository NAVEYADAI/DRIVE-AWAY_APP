package handler

import (
	"back/internal/auth/model"
	"back/internal/auth/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LogIn struct {
	UserName   string `json:"UserName"`
	Password   string `json:"Password"`
	IP         string `json:"IP"`
	MacAddress string `json:"MacAddress"`
}

func LogInUser(c *gin.Context) {
	var logIn LogIn
	err := c.ShouldBindJSON(&logIn)
	if err != nil {
		println(err)
		println("Error in BindJSON")
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "error in BindJSON"}) // 415
		model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 1)
		return
	}
	good, message := service.TryLogInIpAndMAcAddress(logIn.IP, logIn.MacAddress, time.Now())
	if !good {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message}) //409
		model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 2)
		return
	}
	UserId := model.GetUserByUserName(logIn.UserName)
	if UserId == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"}) //404
		model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 3)
		return
	}
	pass := model.GetLastPasswordByUserId(UserId)
	if pass.Pass != logIn.Password {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Password incorrect"}) //401
		model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 4)
		return
	}
	model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, true, time.Now(), -1)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "LogIn Success", "UserId": UserId}) //200
}
