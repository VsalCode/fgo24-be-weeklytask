package routers

import (
	"be-weeklytask/controllers"
	"be-weeklytask/middlewares"

	"github.com/gin-gonic/gin"
)

func usersRouters(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("/profile", controllers.UserProfile)
	r.PUT("/profile", controllers.UpdateUserProfile)
	r.GET("", controllers.ListUsers) 
}
