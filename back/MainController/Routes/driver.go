package routes

import (
	"back/MainController/MainControllerGormModels"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DriverRoute(route *gin.RouterGroup) {
	route.GET("/GetAllDriver", GetAllDriver)
	route.GET("/GetDriverById", GetDriverById)
	route.POST("/CreateDriver", CreateDriver)
	route.PUT("/UpdateDriver", UpdateDriver)
	route.DELETE("/DeleteDriver", DeleteDriver)
}

func DeleteDriver(context *gin.Context) {
	var driver MainControllerGormModels.Driver
	err := context.BindJSON(&driver)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.DeleteDriver(driver)
	context.IndentedJSON(200, gin.H{"message": "Driver deleted"})
}

func UpdateDriver(context *gin.Context) {
	var driver MainControllerGormModels.Driver
	err := context.BindJSON(&driver)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.UpdateDriver(driver)
	context.IndentedJSON(200, gin.H{"message": "Driver updated"})
}

func GetDriverById(context *gin.Context) {
	tmpDriverId := context.Param("id")
	id, err := strconv.Atoi(tmpDriverId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	driver := MainControllerGormModels.GetDriverById(id)
	context.IndentedJSON(http.StatusOK, &driver)
}

func CreateDriver(context *gin.Context) {
	var driver MainControllerGormModels.Driver
	err := context.BindJSON(&driver)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.CreateDriver(driver)
	context.IndentedJSON(200, gin.H{"message": "Driver created"})
}

func GetAllDriver(context *gin.Context) {
	driver := MainControllerGormModels.GetAllDrivers()
	context.IndentedJSON(http.StatusOK, &driver)
}
