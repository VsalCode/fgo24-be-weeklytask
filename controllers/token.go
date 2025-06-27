package controllers

import (
	"be-weeklytask/models"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(ctx *gin.Context) {
	godotenv.Load()

	expirationTime := time.Now().Add(24 * time.Hour)

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": 1,
		"iat":    time.Now().Unix(),
		"exp":    expirationTime.Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("APP_SECRET")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Failed to generate token",
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Token generated successfully",
		Result:  token,
	})
}
