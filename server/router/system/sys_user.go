package system

import (
	v1 "GinVueAdmin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(routeGroup *gin.RouterGroup) {
	userRouter := routeGroup.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi

	{
		userRouter.POST("register", baseApi.Register)
	}
}
