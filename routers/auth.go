package routers

import (
	"github.com/gin-gonic/gin"
	"be-weeklytask/controllers"
)

func authRouters(r *gin.RouterGroup){ 
	r.POST("/register", controllers.AuthRegister)
	r.POST("/login")
	r.GET("/logout")
}