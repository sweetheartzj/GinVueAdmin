package initialize

import (
	"GinVueAdmin/server/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	routerGroup := router.RouterGroupApp

	privateGroup := r.Group("")
	{
		routerGroup.System.InitUserRouter(privateGroup)
	}

	return r
}
