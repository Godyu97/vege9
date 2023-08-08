package vegeRouter

import (
	"github.com/gin-gonic/gin"
)

const (
	uri     = "uri"
	PathUri = "*" + uri
)

type Register interface {
	//200
	SendOk(c *gin.Context, body any)
	//500
	SendBad(c *gin.Context, errMsg string)
	//401
	CheckAuth(c *gin.Context) error
	//map[uri]fnName
	UriToFnName(uri string) string
}

// RegApiHandler
// 根据*uri执行对应方法,req自动从body中取得，其他参数可以自己从ctx中获取
func RegApiHandler(bind Register) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := ctx.Param(uri)
		if len(input) < 2 {
			bind.SendBad(ctx, "afdmWYon uri parsing failed, please req correct path")
			return
		}
		//input为/uri 需要去掉/
		methodName := input[1:]
		methodName = bind.UriToFnName(methodName)
		//权限校验
		if err := bind.CheckAuth(ctx); err != nil {
			return
		}
		//Call Api
		resp, err := Call(ctx, bind, methodName)
		if err != nil {
			bind.SendBad(ctx, err.Error())
			return
		} else {
			bind.SendOk(ctx, resp)
			return
		}
	}
}
