package routers;

import (
	"github.com/gin-gonic/gin"
	"be-weeklytask/controllers"
)

func tokenRouters(r *gin.RouterGroup) {
	r.GET("", controllers.GenerateToken)
}