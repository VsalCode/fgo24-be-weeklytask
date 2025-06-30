package routers;

import (
	"github.com/gin-gonic/gin"
)

func CombineRouters(r *gin.Engine) {
	authRouters(r.Group("/auth"))
	usersRouters(r.Group("/users"))
	profileRouters(r.Group("/profile"))
	transactionRouters(r.Group("/transactions"))
	walletsRouters(r.Group("/wallets"))
}