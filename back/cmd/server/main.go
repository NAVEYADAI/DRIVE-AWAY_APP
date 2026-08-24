package main

import (
	"back/internal/auth"
	"back/internal/config"
	"back/internal/management"
	"back/internal/sorting"

	"regexp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var localhostOrigin = regexp.MustCompile(`^http://localhost:\d+$`)

func main() {
	Server()

}
func Server() {

	config.Connect()
	router := gin.Default()
	// Configure CORS
	router.Use(cors.New(cors.Config{
		// Any localhost port is allowed in dev since the frontend's dev-server
		// port shifts when another local project already holds 3000.
		AllowOriginFunc:  func(origin string) bool { return localhostOrigin.MatchString(origin) },
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
	management.MainController(MAINCONTROLLER)
	router.POST("/Sort", sorting.CreateSort)
	LOGINAPI := router.Group("/loginApi")
	auth.LogIn(LOGINAPI)
	router.POST("/login", auth.Login)
	err := router.Run(":2909")
	if err != nil {
		return
	}

}
