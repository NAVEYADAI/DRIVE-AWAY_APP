package routes

import (
	"back/internal/management/model"
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
	var drive model.Drive
	err := context.ShouldBindJSON(&drive)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.DeleteDrive(drive)
	context.IndentedJSON(200, gin.H{"message": "Drive deleted"})
}

func UpdateDrive(context *gin.Context) {
	var drive model.Drive
	err := context.ShouldBindJSON(&drive)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.UpdateDrive(drive)
	context.IndentedJSON(200, gin.H{"message": "Drive updated"})
}

func CreateDrive(context *gin.Context) {
	var drive model.Drive
	err := context.ShouldBindJSON(&drive)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.CreateDrive(drive)
	context.IndentedJSON(200, gin.H{"message": "Drive created"})
}

func GetDriveById(context *gin.Context) {
	tmpDriveId := context.Param("id")
	id, err := strconv.Atoi(tmpDriveId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	drive := model.GetDriveById(id)
	context.IndentedJSON(http.StatusOK, &drive)
}

func GetAllDrive(context *gin.Context) {
	tmpDrives := model.GetAllDrives()
	context.IndentedJSON(http.StatusOK, &tmpDrives)
}
