package controllers

import (
	"be-weeklytask/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRegister(ctx *gin.Context) {
	var tempData models.User

	err := ctx.ShouldBindJSON(&tempData); 
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Validasi input gagal",
			Error:   err.Error(), 
		})
		return
	}

	if err := models.HandleRegister(tempData); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "User sudah terdaftar atau registrasi gagal",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Registrasi berhasil",
	})
}
