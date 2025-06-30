package routers

import (
	"be-weeklytask/controllers"
	"github.com/gin-gonic/gin"
)

func authRouters(r *gin.RouterGroup){
	r.POST("/register", controllers.AuthRegister)
	r.POST("/login", controllers.AuthLogin)
	// r.GET("/logout")
}