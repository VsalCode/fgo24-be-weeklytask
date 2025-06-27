package controllers

import (
	"be-weeklytask/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Login successful",
		Result:  "Welcome to the application!",
	})

}
