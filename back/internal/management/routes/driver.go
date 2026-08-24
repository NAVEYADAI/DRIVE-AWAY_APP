package routes

import (
	"back/internal/management/model"
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
	var driver model.Driver
	err := context.ShouldBindJSON(&driver)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.DeleteDriver(driver)
	context.IndentedJSON(200, gin.H{"message": "Driver deleted"})
}

func UpdateDriver(context *gin.Context) {
	var driver model.Driver
	err := context.ShouldBindJSON(&driver)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.UpdateDriver(driver)
	context.IndentedJSON(200, gin.H{"message": "Driver updated"})
}

func GetDriverById(context *gin.Context) {
	tmpDriverId := context.Param("id")
	id, err := strconv.Atoi(tmpDriverId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	driver := model.GetDriverById(id)
	context.IndentedJSON(http.StatusOK, &driver)
}

func CreateDriver(context *gin.Context) {
	var driver model.Driver
	err := context.ShouldBindJSON(&driver)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.CreateDriver(driver)
	context.IndentedJSON(200, gin.H{"message": "Driver created"})
}

func GetAllDriver(context *gin.Context) {
	driver := model.GetAllDrivers()
	context.IndentedJSON(http.StatusOK, &driver)
}
