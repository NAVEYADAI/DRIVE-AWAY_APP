package routes

import (
	"back/internal/auth/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func TryLoginRoute(route *gin.RouterGroup) {
	route.GET("/GetAllTryLogin", GetAllTryLogin)
	route.GET("/GetTryLogin/:id", GetTryLogin)
	route.POST("/CreateTryLogin", CreateTryLogin)
	route.PUT("/UpdateTryLogin", UpdateTryLogin)
	route.DELETE("/DeleteTryLogin", DeleteTryLogin)
}

func DeleteTryLogin(context *gin.Context) {
	var tryLogin model.TryLogIn
	err := context.ShouldBindJSON(&tryLogin)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.DeleteTryLogIn(tryLogin)
	context.IndentedJSON(200, gin.H{"message": "TryLogIn deleted"})
}

func UpdateTryLogin(context *gin.Context) {
	var tryLogin model.TryLogIn
	err := context.ShouldBindJSON(&tryLogin)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.UpdateTryLogIn(tryLogin)
	context.IndentedJSON(200, gin.H{"message": "TryLogIn updated"})
}

func CreateTryLogin(context *gin.Context) {
	var tryLogin model.TryLogIn
	err := context.ShouldBindJSON(&tryLogin)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.AddTryLogIn(tryLogin.UserName, tryLogin.Password, tryLogin.IP, tryLogin.MacAddress, tryLogin.Connect, time.Now(), tryLogin.Cause)
	context.IndentedJSON(200, gin.H{"message": "TryLogIn created"})
}

func GetAllTryLogin(context *gin.Context) {
	tryLogin := model.GetAllTryLogIn()
	context.IndentedJSON(http.StatusOK, &tryLogin)
}

func GetTryLogin(context *gin.Context) {
	tmpId := context.Param("id")
	id, err := strconv.Atoi(tmpId)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	tryLogin := model.GetTryLoginById(id)
	context.IndentedJSON(http.StatusOK, &tryLogin)
}
