package routes

import (
	"back/internal/auth/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PersonalInformationRoute(route *gin.RouterGroup) {
	route.GET("/GetPersonalInformation", getPersonalInformation)
	route.GET("/GetPersonalInformation/:id", getPersonalInformationById)
	route.POST("/CreatePersonalInformation", createPersonalInformation)
	route.PUT("/UpdatePersonalInformation", updatePersonalInformation)
	route.DELETE("/DeletePersonalInformation", deletePersonalInformation)
}

func deletePersonalInformation(context *gin.Context) {
	var personalInformation model.PersonalInformation
	err := context.ShouldBindJSON(&personalInformation)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.DeletePersonalInformation(personalInformation)
	context.IndentedJSON(200, gin.H{"message": "PersonalInformation deleted"})
}

func updatePersonalInformation(context *gin.Context) {
	var personalInformation model.PersonalInformation
	err := context.ShouldBindJSON(&personalInformation)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.UpdatePersonalInformation(personalInformation)
	context.IndentedJSON(200, gin.H{"message": "PersonalInformation updated"})
}

func createPersonalInformation(context *gin.Context) {
	var personalInformation model.PersonalInformation
	err := context.ShouldBindJSON(&personalInformation)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.AddPersonalInformation(personalInformation.FName, personalInformation.LName, personalInformation.IdentityCard, personalInformation.UserId)
	context.IndentedJSON(200, gin.H{"message": "PersonalInformation created"})
}

func getPersonalInformationById(context *gin.Context) {
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	personalInformation := model.GetPersonalInformationById(id)
	context.IndentedJSON(200, &personalInformation)
}

func getPersonalInformation(context *gin.Context) {
	personalInformation := model.GetAllPersonalInformation()
	context.IndentedJSON(200, &personalInformation)
}
