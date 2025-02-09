package routes

import (
	"back/MainController/MainControllerGormModels"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CompanyRoute(route *gin.RouterGroup) {
	route.GET("/GetAllCompany", GetAllCompany)
	route.GET("/GetCompanyById", GetCompanyById)
	route.POST("/CreateCompany", CreateCompany)
	route.PUT("/UpdateCompany", UpdateCompany)
	route.DELETE("/DeleteCompany", DeleteCompany)
}

func DeleteCompany(context *gin.Context) {
	var company MainControllerGormModels.Company
	err := context.BindJSON(&company)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.DeleteCompany(company)
	context.IndentedJSON(200, gin.H{"message": "Company deleted"})
}

func UpdateCompany(context *gin.Context) {
	var company MainControllerGormModels.Company
	err := context.BindJSON(&company)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.UpdateCompany(company)
	context.IndentedJSON(200, gin.H{"message": "Company updated"})
}

func CreateCompany(context *gin.Context) {
	var company MainControllerGormModels.Company
	err := context.BindJSON(&company)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	MainControllerGormModels.CreateCompany(company)
	context.IndentedJSON(200, gin.H{"message": "Company created"})
}

func GetCompanyById(context *gin.Context) {
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	company := MainControllerGormModels.GetCompanyById(id)
	context.IndentedJSON(http.StatusOK, &company)
}

func GetAllCompany(context *gin.Context) {
	tmpCompany := MainControllerGormModels.GetAllCompanies()
	context.IndentedJSON(http.StatusOK, &tmpCompany)
}
