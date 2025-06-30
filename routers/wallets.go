package routers

import (
	"be-weeklytask/controllers"
	"be-weeklytask/middlewares"
	"github.com/gin-gonic/gin"
)

func walletsRouters(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("/balance", controllers.Balance)
	r.GET("/records", controllers.FinanceRecords)
}
