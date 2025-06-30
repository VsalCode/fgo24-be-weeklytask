package routers

import (
	"be-weeklytask/controllers"
	"be-weeklytask/middlewares"
	"github.com/gin-gonic/gin"
)

func profileRouters(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("", controllers.UserProfile)
	r.PATCH("", controllers.UpdateUserProfile)
	r.PUT("", controllers.UploadAvatar)
}
