package auth

import (
	"back/internal/auth/handler"
	"back/internal/auth/model"
	"back/internal/auth/routes"
	"back/internal/auth/service"
	"back/internal/config"
	managementmodel "back/internal/management/model"
	"fmt"
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
	err := context.ShouldBindJSON(&logIn)
	if err != nil {
		context.IndentedJSON(415, gin.H{"messege": "error in BindJSON"})
		model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 1)
		return
	}
	good, message := service.TryLogInIpAndMAcAddress(logIn.IP, logIn.MacAddress, time.Now())
	if !good {
		context.IndentedJSON(http.StatusConflict, gin.H{"messege": message})
		model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 2)
		return
	}
	UserId := model.GetUserByUserName(logIn.UserName)
	if UserId == -1 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"messege": "User not found"})
		model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 3)
		return
	}
	pass := model.GetLastPasswordByUserId(UserId)
	if pass.Pass != logIn.Password {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"messege": "Password incorrect"})
		model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, false, time.Now(), 4)
		return
	}
	model.AddTryLogIn(logIn.UserName, logIn.Password, logIn.IP, logIn.MacAddress, true, time.Now(), -1)
	perm := managementmodel.GetPermissionsByUserId(UserId)
	if perm.GetId() != 0 {
		context.IndentedJSON(http.StatusOK, gin.H{"messege": "LogIn Success", "UserId": UserId, "level": perm.GetLevel()})
		return
	}
	context.IndentedJSON(http.StatusUnauthorized, gin.H{"messege": "User has no permissions"})
}

func LogIn(route *gin.RouterGroup) {
	RestApi(route)
	err := config.DB.AutoMigrate(
		&model.DoubleRegistration{},
		&model.Mail{},
		&model.Name{},
		&model.Password{},
		&model.Phone{},
		&model.TryLogIn{},
		&model.Users{},
		&model.PersonalInformation{},
	)
	if err != nil {
		fmt.Printf("Failed to migrate database: %v", err)
	}
	route.POST("/SingUp", handler.CreateUser)
	route.POST("/LogIn", handler.LogInUser)
	route.POST("/ChangeUsername", handler.ChangeUsername)
	route.POST("/ChangePassword", handler.ChangePassword)
	route.POST("/SendMail", handler.SendMail)
	route.POST("/ForgetMyUserName", handler.ForgetMyUserName)
}
func RestApi(route *gin.RouterGroup) {
	MAILAPI := route.Group("/Mail")
	routes.MailRoute(MAILAPI)
	PASSWORDAPI := route.Group("/Password")
	routes.PasswordRoute(PASSWORDAPI)
	PERSONALINFORMATIONAPI := route.Group("/PersonalInformation")
	routes.PersonalInformationRoute(PERSONALINFORMATIONAPI)
	PHONEAPI := route.Group("/Phone")
	routes.PhoneRoute(PHONEAPI)
	TRYLOGINAPI := route.Group("/TryLogin")
	routes.TryLoginRoute(TRYLOGINAPI)
	USERSAPI := route.Group("/Users")
	routes.UserRoute(USERSAPI)
}
