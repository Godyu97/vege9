package vegeRouter

import (
	"github.com/gin-gonic/gin"
	"io"
)

type Register interface {
	//200
	SendOk(c *gin.Context, body any)
	//500
	SendBad(c *gin.Context, errMsg string, body any)
	//401
	CheckAuth(c *gin.Context, body any) error
}

// RegApiHandler
// 根据*uri执行对应方法
func RegApiHandler(bind Register) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Param("uri")
		if len(uri) < 2 {
			bind.SendBad(ctx, "afdmWYon uir解析失败，请访问正确的路径", nil)
			return
		}
		methodName := uri[1:]
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
