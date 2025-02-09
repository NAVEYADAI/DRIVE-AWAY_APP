package routes

import (
	"back/LogInApi/LogInApiGormModels"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.RouterGroup) {
	route.GET("/GetUsers", getUsers)
	route.GET("/GetUser/:id", getUser)
	route.POST("/CreateUser", createUser)
	route.PUT("/UpdateUser", updateUser)
	route.DELETE("/DeleteUser", deleteUser)

}

func deleteUser(context *gin.Context) {
	var user LogInApiGormModels.Users
	err := context.BindJSON(&user)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.DeleteUser(user)
	context.IndentedJSON(200, gin.H{"message": "User deleted"})
}

func updateUser(context *gin.Context) {
	var user LogInApiGormModels.Users
	err := context.BindJSON(&user)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.UpdateUser(user.UserName, user.ID)
	context.IndentedJSON(200, gin.H{"message": "User updated"})

}

func createUser(context *gin.Context) {
	var user LogInApiGormModels.Users
	err := context.BindJSON(&user)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "error in BindJSON"})
		return
	}
	LogInApiGormModels.AddUser(user.UserName)
}

func getUser(context *gin.Context) {
	var user LogInApiGormModels.Users
	tmpID := context.Param("id")
	id, err := strconv.Atoi(tmpID)
	if err != nil {
		context.IndentedJSON(400, gin.H{"message": "id is not a number"})
		return
	}
	user = LogInApiGormModels.GetUserById(id)
	if user.ID == -1 {
		context.IndentedJSON(404, gin.H{"message": "user not found"})
		return
	}
	context.IndentedJSON(200, &user)
}

func getUsers(context *gin.Context) {
	users := LogInApiGormModels.GetAllUser()
	context.IndentedJSON(200, &users)
}
