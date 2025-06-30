package routers;

import (
	"github.com/gin-gonic/gin"
	"be-weeklytask/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CombineRouters(r *gin.Engine) {
	authRouters(r.Group("/auth"))
	usersRouters(r.Group("/users"))
	profileRouters(r.Group("/profile"))
	transactionRouters(r.Group("/transactions"))
	walletsRouters(r.Group("/wallets"))
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}