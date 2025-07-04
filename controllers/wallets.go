package controllers;

import (
	"be-weeklytask/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Balance(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized!",
		})
		return
	}

	balance, err := models.GetBalance(userId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to get balance",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Get Balance Successfully!",
		Result:  balance,
	})
}

func FinanceRecords(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized!",
		})
		return
	}

	records, err := models.GetFinanceRecords(userId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to get finance records",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Get Finance Records Successfully!",
		Result:  records,
	})
}