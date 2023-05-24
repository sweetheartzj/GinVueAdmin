package system

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(routeGroup *gin.RouterGroup) {
	userRouter := routeGroup.Group("user")

	{
		userRouter.POST("login")
	}
}
