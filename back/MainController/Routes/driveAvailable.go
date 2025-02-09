package routes

import (
	"back/MainController/MainControllerGormModels"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DriveAvailableRoute(route *gin.RouterGroup) {
	route.GET("/GetAllDriveAvailable", GetAllDriveAvailable)
	route.GET("/GetDriveAvailableById", GetDriveAvailableById)
	route.POST("/CreateDriveAvailable", CreateDriveAvailable)
	route.PUT("/UpdateDriveAvailable", UpdateDriveAvailable)
	route.DELETE("/DeleteDriveAvailable", DeleteDriveAvailable)
}

func DeleteDriveAvailable(context *gin.Context) {
	var driverAvailable MainControllerGormModels.DriverAvailable
	err := context.BindJSON(&driverAvailable)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.DeleteDriverAvailable(driverAvailable)
	context.IndentedJSON(200, gin.H{"message": "DriverAvailable deleted"})
}

func UpdateDriveAvailable(context *gin.Context) {
	var driverAvailable MainControllerGormModels.DriverAvailable
	err := context.BindJSON(&driverAvailable)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.UpdateDriverAvailable(driverAvailable)
	context.IndentedJSON(200, gin.H{"message": "DriverAvailable updated"})
}

func GetDriveAvailableById(context *gin.Context) {
	var tmpId = context.Param("id")
	id, err := strconv.Atoi(tmpId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	driverAvailable := MainControllerGormModels.GetDriveAvailableById(id)
	context.IndentedJSON(http.StatusOK, &driverAvailable)

}

func CreateDriveAvailable(context *gin.Context) {
	var driverAvailable MainControllerGormModels.DriverAvailable
	err := context.BindJSON(&driverAvailable)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.CreateDriverAvailable(driverAvailable)
	context.IndentedJSON(200, gin.H{"message": "DriverAvailable created"})
}

func GetAllDriveAvailable(context *gin.Context) {
	driverAvailable := MainControllerGormModels.GetAllDriveAvailables()
	context.IndentedJSON(http.StatusOK, &driverAvailable)
}
