package routes

import (
	"back/internal/management/model"
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
	var company model.Company
	err := context.ShouldBindJSON(&company)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.DeleteCompany(company)
	context.IndentedJSON(200, gin.H{"message": "Company deleted"})
}

func UpdateCompany(context *gin.Context) {
	var company model.Company
	err := context.ShouldBindJSON(&company)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.UpdateCompany(company)
	context.IndentedJSON(200, gin.H{"message": "Company updated"})
}

func CreateCompany(context *gin.Context) {
	var company model.Company
	err := context.ShouldBindJSON(&company)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	model.CreateCompany(company)
	context.IndentedJSON(200, gin.H{"message": "Company created"})
}

func GetCompanyById(context *gin.Context) {
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	company := model.GetCompanyById(id)
	context.IndentedJSON(http.StatusOK, &company)
}

func GetAllCompany(context *gin.Context) {
	tmpCompany := model.GetAllCompanies()
	context.IndentedJSON(http.StatusOK, &tmpCompany)
}
