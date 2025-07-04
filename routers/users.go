package routers

import (
	"be-weeklytask/controllers"
	"be-weeklytask/middlewares"

	"github.com/gin-gonic/gin"
)

func usersRouters(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("", controllers.ListUsers) 
}
