package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sassatricolado/auth-go/database"
)

func main() {
	database.InitDB()

	router := gin.Default()
	router.POST("/signin")
	router.POST("/login")
	router.GET("/home")

	router.Run(":8080")
}
