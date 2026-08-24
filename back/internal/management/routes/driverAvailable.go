package routes

import (
	"back/internal/management/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DriverAvailableRoute(route *gin.RouterGroup) {
	route.GET("/GetAllDriverAvailable", GetAllDriverAvailable)
	route.GET("/GetDriverAvailableById", GetDriverAvailableById)
	route.POST("/CreateDriverAvailable", CreateDriverAvailable)
	route.PUT("/UpdateDriverAvailable", UpdateDriverAvailable)
	route.DELETE("/DeleteDriverAvailable", DeleteDriverAvailable)
}

func UpdateDriverAvailable(context *gin.Context) {
	var driverAvailable model.DriverAvailable
	err := context.ShouldBindJSON(&driverAvailable)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.UpdateDriverAvailable(driverAvailable)
	context.IndentedJSON(200, gin.H{"message": "DriverAvailable updated"})
}

func DeleteDriverAvailable(context *gin.Context) {
	var DriverAvailable model.DriverAvailable
	err := context.ShouldBindJSON(&DriverAvailable)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.DeleteDriverAvailable(DriverAvailable)
	context.IndentedJSON(200, gin.H{"message": "DriverAvailable deleted"})
}

func CreateDriverAvailable(context *gin.Context) {
	var driverAvailable model.DriverAvailable
	err := context.ShouldBindJSON(&driverAvailable)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.CreateDriverAvailable(driverAvailable)
	context.IndentedJSON(200, gin.H{"message": "DriverAvailable created"})
}

func GetDriverAvailableById(context *gin.Context) {
	tmpDriverAvailableId := context.Param("id")
	id, err := strconv.Atoi(tmpDriverAvailableId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	driverAvailable := model.GetDriverAvailableById(id)
	context.IndentedJSON(http.StatusOK, &driverAvailable)
}

func GetAllDriverAvailable(context *gin.Context) {
	driverAvailable := model.GetAllDriverAvailables()
	context.IndentedJSON(http.StatusOK, &driverAvailable)
}
