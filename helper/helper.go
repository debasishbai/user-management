package helper

import (
	"net/http"
	"user-management/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HasValidBody(context *gin.Context, userModel *models.User) bool {
	validate := validator.New()
	if err := validate.Struct(userModel); err != nil {
		context.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return false
	}
	return true
}
