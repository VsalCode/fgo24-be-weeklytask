package controllers

import (
	"be-weeklytask/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRegister(ctx *gin.Context) {
	var tempData models.User

	err := ctx.ShouldBindJSON(&tempData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Validasi invalid!",
			Error:   err.Error(),
		})
		return
	}

	userId, err := models.HandleRegister(tempData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "User already Registered!",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Registrasi Successfully!",
		Result:  userId,
	})
}
