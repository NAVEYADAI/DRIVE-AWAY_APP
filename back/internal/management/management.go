package management

import (
	"back/internal/config"
	"back/internal/management/handler"
	"back/internal/management/handler/driver"
	"back/internal/management/handler/manager"
	"back/internal/management/handler/systemadmin"
	"back/internal/management/model"
	"back/internal/management/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

// localhost:2909/MainController

func MainController(route *gin.RouterGroup) {
	restApi(route)
	err := config.DB.AutoMigrate(
		&model.Company{},
		&model.Drive{},
		&model.DriveAvailable{},
		&model.DriveOfDriver{},
		&model.Driver{},
		&model.DriverAvailable{},
		&model.Permissions{},
		&model.Sort{},

	)
	if err != nil {
		fmt.Printf("Failed to migrate database: %v", err)
	}
	route.POST("/LogIn", handler.LogIn)
	route.POST("/AddCompany", systemadmin.Add_Company)
	route.POST("/AddPosition", systemadmin.Add_Position)
	route.POST("/AddDriver", manager.Add_Driver)
	route.POST("/AddDrive", manager.Add_Drive)
	route.POST("/DeleteDrive", manager.Delete_Drive)
	route.POST("/GetDrives", manager.Get_Drives)
	route.POST("/DeleteDriver", manager.Delete_Driver)
	route.POST("/GetDrivers", manager.Get_Drivers)
	route.POST("/ViewOrNotDrive", manager.ViewOrNotDrive)
	route.POST("/ViewOrNotDriver", manager.ViewOrNotDriver)
	route.POST("/GetMyDrives", driver.GetMyDrives)

	//route.GET("/LogIn", handler.Login)

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
