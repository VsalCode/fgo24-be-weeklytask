package controllers

import (
	"be-weeklytask/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ListUsers(ctx *gin.Context) {
	key := ctx.Query("search")

	result, err := models.FindUserByNameOrPhone(key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to get people",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Users retrieved successfully",
		Result:  result,
	})
}
