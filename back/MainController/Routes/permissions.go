package routes

import (
	"back/MainController/MainControllerGormModels"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PermissionRoute(route *gin.RouterGroup) {
	route.GET("/GetAllPermission", GetAllPermission)
	route.GET("/GetPermissionById", GetPermissionById)
	route.POST("/CreatePermission", CreatePermission)
	route.PUT("/UpdatePermission", UpdatePermission)
	route.DELETE("/DeletePermission", DeletePermission)
}

func DeletePermission(context *gin.Context) {
	var permission MainControllerGormModels.Permissions
	err := context.BindJSON(&permission)
	if err != nil {
		context.IndentedJSON(400, gin.H{"messeage": "error in binding data"})
		return
	}
	MainControllerGormModels.DeletePermissions(permission)
	context.IndentedJSON(http.StatusOK, gin.H{"messeage": "Permission Deleted"})
}

func UpdatePermission(context *gin.Context) {
	var permission MainControllerGormModels.Permissions
	err := context.BindJSON(&permission)
	if err != nil {
		context.IndentedJSON(400, gin.H{"messeage": "error in binding data"})
		return
	}
	MainControllerGormModels.UpdatePermissions(permission)
	context.IndentedJSON(http.StatusOK, gin.H{"messeage": "Permission Updated"})
}

func CreatePermission(context *gin.Context) {
	var permission MainControllerGormModels.Permissions
	err := context.BindJSON(&permission)
	if err != nil {
		context.IndentedJSON(400, gin.H{"messeage": "error in binding data"})
		return
	}
	MainControllerGormModels.CreatePermissions(permission)
	context.IndentedJSON(http.StatusOK, gin.H{"messeage": "Permission Created"})
}

func GetAllPermission(context *gin.Context) {
	permission := MainControllerGormModels.GetAllPermissions()
	context.IndentedJSON(http.StatusOK, &permission)
}

func GetPermissionById(context *gin.Context) {
	tmpId := context.Param("id")
	id, err := strconv.Atoi(tmpId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"messeage": "id is not a number"})
		return
	}
	permission := MainControllerGormModels.GetPermissionsById(id)
	context.IndentedJSON(http.StatusOK, permission)
}
