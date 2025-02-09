package routes

import (
	"back/MainController/MainControllerGormModels"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DriveRoute(route *gin.RouterGroup) {
	route.GET("/GetAllDrive", GetAllDrive)
	route.GET("/GetDriveById", GetDriveById)
	route.POST("/CreateDrive", CreateDrive)
	route.PUT("/UpdateDrive", UpdateDrive)
	route.DELETE("/DeleteDrive", DeleteDrive)
}

func DeleteDrive(context *gin.Context) {
	var drive MainControllerGormModels.Drive
	err := context.BindJSON(&drive)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.DeleteDrive(drive)
	context.IndentedJSON(200, gin.H{"message": "Drive deleted"})
}

func UpdateDrive(context *gin.Context) {
	var drive MainControllerGormModels.Drive
	err := context.BindJSON(&drive)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.UpdateDrive(drive)
	context.IndentedJSON(200, gin.H{"message": "Drive updated"})
}

func CreateDrive(context *gin.Context) {
	var drive MainControllerGormModels.Drive
	err := context.BindJSON(&drive)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.CreateDrive(drive)
	context.IndentedJSON(200, gin.H{"message": "Drive created"})
}

func GetDriveById(context *gin.Context) {
	tmpDriveId := context.Param("id")
	id, err := strconv.Atoi(tmpDriveId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	drive := MainControllerGormModels.GetDriveById(id)
	context.IndentedJSON(http.StatusOK, &drive)
}

func GetAllDrive(context *gin.Context) {
	tmpDrives := MainControllerGormModels.GetAllDrives()
	context.IndentedJSON(http.StatusOK, &tmpDrives)
}
