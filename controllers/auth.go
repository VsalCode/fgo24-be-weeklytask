package controllers

import (
	"be-weeklytask/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func AuthLogin(ctx *gin.Context) {
	godotenv.Load()

	loginData := models.LoginRequest{}
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request payload",
		})
		return
	}

	user, err := models.FindUserByEmail(loginData.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Invalid email or password!",
		})
		return
	}

	if user.Password != loginData.Password {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Invalid email or password!",
		})
		return
	}

	token, err := generateToken(user.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Login successful",
		Result:  token,
	})

}
