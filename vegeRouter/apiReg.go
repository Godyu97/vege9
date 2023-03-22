package vegeRouter

import (
	"github.com/gin-gonic/gin"
	"io"
)

type Apier interface {
	//200
	SendOk(c *gin.Context, body any)
	//500
	SendBad(c *gin.Context, message string, body any)
	//401
	CheckAuth(c *gin.Context, body any) error
}

// RegApiHandler
// 根据*uri执行对应方法
func RegApiHandler(bind Apier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		methodName := ctx.Param("uri")[1:]
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			bind.SendBad(ctx, err.Error(), nil)
			return
		}
		if err = bind.CheckAuth(ctx, nil); err != nil {
			return
		}
		resp, err := Call(ctx, bind, methodName, string(body))
		if err != nil {
			bind.SendBad(ctx, err.Error(), nil)
			return
		} else {
			bind.SendOk(ctx, resp)
			return
		}
	}
}
