package routers

import (
	"be-weeklytask/controllers"
	"be-weeklytask/middlewares"

	"github.com/gin-gonic/gin"
)

func transactionRouters(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.POST("/topup", controllers.Topup)
	r.POST("/transfer", controllers.Transfer)
	r.GET("/history", controllers.HistoryTransactions )
}
