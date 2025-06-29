package routers;

import (
	"github.com/gin-gonic/gin"
)

func CombineRouters(r *gin.Engine) {
	authRouters(r.Group("/auth"))
	usersRouters(r.Group("/users"))
	transactionRouters(r.Group("/transactions"))
	balanceRouters(r.Group("/balance"))
}