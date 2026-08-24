package routes

import (
	"back/internal/management/model"
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
	var permission model.Permissions
	err := context.ShouldBindJSON(&permission)
	if err != nil {
		context.IndentedJSON(400, gin.H{"messeage": "error in binding data"})
		return
	}
	model.DeletePermissions(permission)
	context.IndentedJSON(http.StatusOK, gin.H{"messeage": "Permission Deleted"})
}

func UpdatePermission(context *gin.Context) {
	var permission model.Permissions
	err := context.ShouldBindJSON(&permission)
	if err != nil {
		context.IndentedJSON(400, gin.H{"messeage": "error in binding data"})
		return
	}
	model.UpdatePermissions(permission)
	context.IndentedJSON(http.StatusOK, gin.H{"messeage": "Permission Updated"})
}

func CreatePermission(context *gin.Context) {
	var permission model.Permissions
	err := context.ShouldBindJSON(&permission)
	if err != nil {
		context.IndentedJSON(400, gin.H{"messeage": "error in binding data"})
		return
	}
	model.CreatePermissions(permission)
	context.IndentedJSON(http.StatusOK, gin.H{"messeage": "Permission Created"})
}

func GetAllPermission(context *gin.Context) {
	permission := model.GetAllPermissions()
	context.IndentedJSON(http.StatusOK, &permission)
}

func GetPermissionById(context *gin.Context) {
	tmpId := context.Param("id")
	id, err := strconv.Atoi(tmpId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"messeage": "id is not a number"})
		return
	}
	permission := model.GetPermissionsById(id)
	context.IndentedJSON(http.StatusOK, permission)
}
