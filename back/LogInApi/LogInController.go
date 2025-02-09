package LogInApi

import (
	"back/LogInApi/LogInApiController"
	"back/LogInApi/config"
	"back/LogInApi/routes"
	"fmt"
	"back/LogInApi/LogInApiGormModels"

	"github.com/gin-gonic/gin"
)

func LogIn(route *gin.RouterGroup) {
	RestApi(route)
	err := config.DB.AutoMigrate(
		&LogInApiGormModels.DoubleRegistration{},
		&LogInApiGormModels.Mail{},
		&LogInApiGormModels.Name{},
		&LogInApiGormModels.Password{},
		&LogInApiGormModels.Phone{},
		&LogInApiGormModels.TryLogIn{},
		&LogInApiGormModels.Users{},
		&LogInApiGormModels.PersonalInformation{},
	)
	if err != nil {
		fmt.Print("Failed to migrate database: %v", err)
	}
	route.POST("/SingUp", LogInApiController.CreateUser)
	route.POST("/LogIn", LogInApiController.LogInUser)
	route.POST("/ChangeUsername", LogInApiController.ChangeUsername)
	route.POST("/ChangePassword", LogInApiController.ChangePassword)
	route.POST("/SendMail", LogInApiController.SendMail)
	route.POST("/ForgetMyUserName", LogInApiController.ForgetMyUserName)
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
func UpDbLogIn() {
	config.UpDBLogIn()
}
