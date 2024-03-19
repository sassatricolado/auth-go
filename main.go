package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/signin")
	router.POST("/login")
	router.GET("/home")

	router.Run(":8080")
}
