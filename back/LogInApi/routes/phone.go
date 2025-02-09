package routes

import (
	"back/LogInApi/LogInApiGormModels"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PhoneRoute(route *gin.RouterGroup) {
	route.GET("/GetAllPhone", GetAllPhone)
	route.GET("/GetPhoneById/:id", GetAllPhoneById)
	route.GET("/GetPhoneByUserId/:id", GetAllPhoneByUserId)
	route.POST("/CreatePhone", CreatePhone)
	route.PUT("/UpdatePhone", UpdatePhone)
	route.DELETE("/DeletePhone", DeletePhone)
}

func DeletePhone(context *gin.Context) {
	var phone LogInApiGormModels.Phone
	err := context.BindJSON(&phone)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.DeletePhone(phone)
	context.IndentedJSON(200, gin.H{"message": "Phone deleted"})
}

func UpdatePhone(context *gin.Context) {
	var phone LogInApiGormModels.Phone
	err := context.BindJSON(&phone)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.UpdatePhone(phone)
	context.IndentedJSON(200, gin.H{"message": "Phone updated"})
}

func CreatePhone(context *gin.Context) {
	var phone LogInApiGormModels.Phone
	err := context.BindJSON(&phone)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.AddPhone(phone.Phone, phone.UserId)
	context.IndentedJSON(200, gin.H{"message": "Phone created"})
}

func GetAllPhoneByUserId(context *gin.Context) {
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	phones := LogInApiGormModels.GetPhoneByUserId(id)
	context.IndentedJSON(200, &phones)
}

func GetAllPhoneById(context *gin.Context) {
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	phone := LogInApiGormModels.GetPhoneById(id)
	context.IndentedJSON(200, &phone)
}

func GetAllPhone(context *gin.Context) {
	tmpPhones := LogInApiGormModels.GetAllPhone()
	context.IndentedJSON(200, &tmpPhones)
}
