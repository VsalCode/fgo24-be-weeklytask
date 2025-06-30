package controllers

import (
	"be-weeklytask/dto"
	"be-weeklytask/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "User registration data"
// @Success 200 {object} models.Response{result=int} "User ID"
// @Failure 400 {object} models.Response "Validation error or already registered"
// @Router /auth/register [post]
func AuthRegister(ctx *gin.Context) {
	var tempData dto.RegisterRequest

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

// @Summary Login a user
// @Description Login a user with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "User login data"
// @Success 200 {object} models.Response{result=string} "JWT token"
// @Failure 400 {object} models.Response "Invalid request payload"
// @Failure 401 {object} models.Response "Invalid email or password"
// @Router /auth/login [post]
func AuthLogin(ctx *gin.Context) {
	godotenv.Load()

	loginData := dto.LoginRequest{}
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
			Message: "Invalid email or pin!",
		})
		return
	}

	if user.Password != loginData.Password {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Invalid password!",
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
