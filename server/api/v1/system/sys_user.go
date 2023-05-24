package system

import (
	commonRes "GinVueAdmin/server/model/common/response"
	systemReq "GinVueAdmin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (base *BaseApi) Login(ctx *gin.Context) {
	var l systemReq.Login
	_ = ctx.ShouldBindJSON(&l)
}

func (base *BaseApi) Register(ctx *gin.Context) {
	var r systemReq.Register
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		commonRes.FailWithMessage("Parameters Error", ctx)
	}
}
