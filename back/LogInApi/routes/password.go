package routes

import (
	"back/LogInApi/LogInApiGormModels"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PasswordRoute(route *gin.RouterGroup) {
	route.GET("/GetAllPassword", GetAllPassword)
	route.GET("/GetPasswordById/:id", GetAllPasswordById)
	route.GET("/GetPasswordByUserId/:id", GetAllPasswordByUserId)
	route.POST("/CreatePassword", CreatePassword)
	route.PUT("/UpdatePassword", UpdatePassword)
	route.DELETE("/DeletePassword", DeletePassword)
}

func DeletePassword(context *gin.Context) {
	var password LogInApiGormModels.Password
	err := context.BindJSON(&password)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.DeletePassword(password)
	context.IndentedJSON(200, gin.H{"message": "Password deleted"})
}

func UpdatePassword(context *gin.Context) {
	var password LogInApiGormModels.Password
	err := context.BindJSON(&password)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.UpdatePassword(password)
	context.IndentedJSON(200, gin.H{"message": "Password updated"})
}

func GetAllPasswordById(context *gin.Context) {
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	password := LogInApiGormModels.GetPasswordById(id)
	context.IndentedJSON(http.StatusOK, &password)
}

func GetAllPasswordByUserId(context *gin.Context) {
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	passwords := LogInApiGormModels.GetPasswordByUserId(id)
	context.IndentedJSON(http.StatusOK, &passwords)
}

func GetAllPassword(context *gin.Context) {
	passwords := LogInApiGormModels.GetAllPassword()
	context.IndentedJSON(http.StatusOK, &passwords)
}

func CreatePassword(context *gin.Context) {
	var password LogInApiGormModels.Password
	err := context.BindJSON(&password)
	if err != nil {
		println(err)
		println("Error in BindJSON")
		context.IndentedJSON(http.StatusUnsupportedMediaType,
			gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.AddPassword(password.Pass, password.UserId)
}
