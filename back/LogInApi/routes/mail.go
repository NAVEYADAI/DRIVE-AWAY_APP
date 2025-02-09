package routes

import (
	"back/LogInApi/LogInApiGormModels"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MailRoute(route *gin.RouterGroup) {
	route.GET("/GetAllMail", GetAllMail)
	route.GET("/GetMail/:id", GetMail)
	route.GET("/GetMailByUserId/:id", GetMailByUserId)
	route.POST("/CreateMail", CreateMail)
	route.PUT("/UpdateMail", UpdateMail)
	route.DELETE("/DeleteMail", DeleteMail)
}

func DeleteMail(context *gin.Context) {
	var mail LogInApiGormModels.Mail
	err := context.BindJSON(&mail)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.DeleteMail(mail)
	context.IndentedJSON(200, gin.H{"message": "Mail deleted"})
}

func UpdateMail(context *gin.Context) {
	var mail LogInApiGormModels.Mail
	err := context.BindJSON(&mail)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.UpdateMail(mail)
	context.IndentedJSON(200, gin.H{"message": "Mail updated"})
}

func CreateMail(context *gin.Context) {
	var mail LogInApiGormModels.Mail
	err := context.BindJSON(&mail)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.AddMail(mail.Mail, mail.UserId)
	context.IndentedJSON(200, gin.H{"message": "Mail created"})
}

func GetMailByUserId(context *gin.Context) {
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	mail := LogInApiGormModels.GetMailByUserId(id)
	if mail.ID == -1 {
		context.IndentedJSON(404, gin.H{"message": "mail not found"})
		return
	}
	context.IndentedJSON(200, &mail)
}

func GetMail(context *gin.Context) {
	var mail LogInApiGormModels.Mail
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	mail = LogInApiGormModels.GetMailById(id)
	if mail.ID == -1 {
		context.IndentedJSON(404, gin.H{"message": "mail not found"})
		return
	}
	context.IndentedJSON(200, &mail)
}
func GetAllMail(c *gin.Context) {
	ListOfMail := LogInApiGormModels.GetAllMail()
	c.IndentedJSON(http.StatusOK, &ListOfMail)
}
