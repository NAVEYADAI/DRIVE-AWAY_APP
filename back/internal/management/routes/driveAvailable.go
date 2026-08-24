package routes

import (
	"back/internal/management/model"
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
	var driverAvailable model.DriverAvailable
	err := context.ShouldBindJSON(&driverAvailable)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.DeleteDriverAvailable(driverAvailable)
	context.IndentedJSON(200, gin.H{"message": "DriverAvailable deleted"})
}

func UpdateDriveAvailable(context *gin.Context) {
	var driverAvailable model.DriverAvailable
	err := context.ShouldBindJSON(&driverAvailable)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.UpdateDriverAvailable(driverAvailable)
	context.IndentedJSON(200, gin.H{"message": "DriverAvailable updated"})
}

func GetDriveAvailableById(context *gin.Context) {
	var tmpId = context.Param("id")
	id, err := strconv.Atoi(tmpId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	driverAvailable := model.GetDriveAvailableById(id)
	context.IndentedJSON(http.StatusOK, &driverAvailable)

}

func CreateDriveAvailable(context *gin.Context) {
	var driverAvailable model.DriverAvailable
	err := context.ShouldBindJSON(&driverAvailable)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.CreateDriverAvailable(driverAvailable)
	context.IndentedJSON(200, gin.H{"message": "DriverAvailable created"})
}

func GetAllDriveAvailable(context *gin.Context) {
	driverAvailable := model.GetAllDriveAvailables()
	context.IndentedJSON(http.StatusOK, &driverAvailable)
}
