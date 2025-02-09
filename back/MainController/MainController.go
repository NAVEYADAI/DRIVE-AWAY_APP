package MainController

import (
	routes "back/MainController/Routes"
	"back/MainController/config"
	"back/MainController/controller"
	"back/MainController/controller/Driver"
	"back/MainController/controller/Manager"
	"back/MainController/controller/SystemAdministrator"
	"back/MainController/MainControllerGormModels"
	"fmt"

	"github.com/gin-gonic/gin"
)

// localhost:2909/MainController

func MainController(route *gin.RouterGroup) {
	restApi(route)
	err := config.DB.AutoMigrate(
		&MainControllerGormModels.Company{},
		&MainControllerGormModels.Drive{},
		&MainControllerGormModels.DriveAvailable{},
		&MainControllerGormModels.DriveOfDriver{},
		&MainControllerGormModels.Driver{},
		&MainControllerGormModels.DriverAvailable{},
		&MainControllerGormModels.Permissions{},
		&MainControllerGormModels.Sort{},

	)
	if err != nil {
		fmt.Print("Failed to migrate database: %v", err)
	}
	route.POST("/LogIn", controller.LogIn)
	route.POST("/AddCompany", SystemAdministrator.Add_Company)
	route.POST("/AddPosition", SystemAdministrator.Add_Position)
	route.POST("/AddDriver", Manager.Add_Driver)
	route.POST("/AddDrive", Manager.Add_Drive)
	route.POST("/DeleteDrive", Manager.Delete_Drive)
	route.POST("/GetDrives", Manager.Get_Drives)
	route.POST("/DeleteDriver", Manager.Delete_Driver)
	route.POST("/GetDrivers", Manager.Get_Drivers)
	route.POST("/ViewOrNotDrive", Manager.ViewOrNotDrive)
	route.POST("/ViewOrNotDriver", Manager.ViewOrNotDriver)
	route.POST("/GetMyDrives", Driver.GetMyDrives)

	//route.GET("/LogIn", controller.Login)

}
func restApi(route *gin.RouterGroup) {
	COMPANYAPI := route.Group("/Company")
	routes.CompanyRoute(COMPANYAPI)
	DRIVEAPI := route.Group("/Drive")
	routes.DriveRoute(DRIVEAPI)
	DRIVEAVAILABLEAPI := route.Group("/DriveAvailable")
	routes.DriveAvailableRoute(DRIVEAVAILABLEAPI)
	DRIVERAPI := route.Group("/Driver")
	routes.DriverRoute(DRIVERAPI)
	DRIVERAVAILABLEAPI := route.Group("/DriverAvailable")
	routes.DriverAvailableRoute(DRIVERAVAILABLEAPI)
	PERMISSIONSAPI := route.Group("/Permissions")
	routes.PermissionRoute(PERMISSIONSAPI)
}
func UpDbMainController() {
	config.UpDBMainController()
}
