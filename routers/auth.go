package routers

import (
	"be-weeklytask/controllers"
	// "be-weeklytask/middlewares"
	"github.com/gin-gonic/gin"
)

func authRouters(r *gin.RouterGroup){
	r.POST("/register", controllers.AuthRegister)
	// r.Use(middlewares.VerifyToken()) 
	r.POST("/login", controllers.AuthLogin)
	r.GET("/logout")
}