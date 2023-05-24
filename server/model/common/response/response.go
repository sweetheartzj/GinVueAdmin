package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0
	ERROR   = 7
)

func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &Response{
		code,
		data,
		msg,
	})
	return
}

func Ok(data interface{}, msg string, ctx *gin.Context) {
	Result(SUCCESS, data, msg, ctx)
	return
}

func OKInDefault(ctx *gin.Context) {
	Result(SUCCESS, nil, "success", ctx)
	return
}

func OkWithMessage(msg string, ctx *gin.Context) {
	Result(SUCCESS, nil, msg, ctx)
}

func OkWithData(data interface{}, ctx *gin.Context) {
	Result(SUCCESS, data, "success", ctx)
	return
}

func Fail(data interface{}, msg string, ctx *gin.Context) {
	Result(ERROR, data, msg, ctx)
	return
}

func FailInDefault(ctx *gin.Context) {
	Result(ERROR, nil, "error", ctx)
	return
}

func FailWithMessage(msg string, ctx *gin.Context) {
	Result(ERROR, nil, msg, ctx)
	return
}

func FailWithData(data interface{}, ctx *gin.Context) {
	Result(ERROR, data, "error", ctx)
	return
}
