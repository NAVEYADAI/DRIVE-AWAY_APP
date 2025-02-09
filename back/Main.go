package main

import (
	"back/LogInApi"
	"back/MainController"
	"back/MainSorting/MainSort"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	Server()

}
func Server() {

	LogInApi.UpDbLogIn()
	MainController.UpDbMainController()
	router := gin.Default()
	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running"})
	})
	MAINCONTROLLER := router.Group("/MainController")
	MainController.MainController(MAINCONTROLLER)
	router.POST("/Sort", MainSort.CreateSort)
	LOGINAPI := router.Group("/loginApi")
	LogInApi.LogIn(LOGINAPI)
	router.POST("/login", LogInApi.Login)
	err := router.Run(":2909")
	if err != nil {
		return
	}

}
