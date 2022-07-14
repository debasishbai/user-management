package controller

import (
	"net/http"
	"user-management/database"
	"user-management/helper"
	"user-management/models"

	"github.com/gin-gonic/gin"
)

// GET user
func GetUser(context *gin.Context) {
	param := context.Query("id")
	user := database.GetUser(param)
	context.JSON(http.StatusOK, user)
}

// CREATE user
func CreateUser(context *gin.Context) {
	userModel := new(models.User)
	if err := context.Bind(userModel); err != nil {
		context.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if !helper.HasValidBody(context, userModel) {
		return
	}
	response, err := database.CreateUser(userModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}
	context.JSON(http.StatusOK, models.Response{Body: response})
}

//UPDATE user
func UpdateUser(context *gin.Context) {
	userID := context.Param("id")
	if userID == "" {
		context.JSON(http.StatusBadRequest, models.Response{Message: "Invalid ID"})
		return
	}
	userModel := new(models.User)
	if err := context.Bind(userModel); err != nil {
		context.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}
	if !helper.HasValidBody(context, userModel) {
		return
	}
	if err := database.UpdateUser(context, userID, userModel); err != nil {
		context.JSON(http.StatusNotFound, models.Response{Message: err.Error()})
		return
	}
	context.JSON(http.StatusOK, models.Response{Message: "Success"})
}

//DELETE user
func DeleteUser(context *gin.Context) {
	userID := context.Param("id")
	if userID == "" {
		context.JSON(http.StatusBadRequest, models.Response{Message: "Invalid ID"})
		return
	}
	err := database.DeleteUser(userID)
	if err != nil {
		context.JSON(http.StatusNotFound, models.Response{Message: err.Error()})
		return
	}
	context.JSON(http.StatusNoContent, models.Response{Message: "Success"})
}
