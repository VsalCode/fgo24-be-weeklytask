package controllers

import (
	"be-weeklytask/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserProfile(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	user, err := models.FindUserById(userId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to get user profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Get User Successfully!",
		Result: user,
	})
}

func UpdateUserProfile(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request payload",
			Error:   err.Error(),
		})
		return
	}

	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	user, err := models.GetUpdateUser(userId.(int), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to update profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Profile updated successfully",
		Result: user,
	})
}
