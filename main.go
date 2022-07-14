package main

import (
	"user-management/controller"
	"user-management/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userManagement := router.Group("/user_management")
	userManagement.GET("/user", controller.GetUser)
	userManagement.POST("/user", controller.CreateUser)
	userManagement.PATCH("/user/:id", controller.UpdateUser)
	userManagement.DELETE("/user/:id", controller.DeleteUser)
	database.ConnectDatabase()
	router.Run()

}
