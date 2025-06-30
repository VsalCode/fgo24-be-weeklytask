package controllers

import (
	"be-weeklytask/dto"
	"be-weeklytask/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
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
		Result:  user,
	})
}

func UpdateUserProfile(ctx *gin.Context) {
	var user dto.UpdatedUser
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
		Result:  user,
	})
}

func UploadAvatar(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}
	userID := userIdRaw.(int)

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Upload failed",
		})
		return
	}

	if file.Size > 2*1024*1024 {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "File is too large",
		})
		return
	}

	filename := fmt.Sprintf("user_%d_%s", userID, file.Filename)
	filepath := "./uploads/" + filename

	ctx.SaveUploadedFile(file, filepath); 
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to save file",
		})
		return
	}
	
	err = models.AddAvatar(userID, filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to update user profile",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Profile picture updated",
		Result: filename,
	})
}