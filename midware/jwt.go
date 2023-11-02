package midware

import (
	"errors"
	"github.com/Godyu97/vege9/midware/jwtApi"
	"github.com/gin-gonic/gin"
	"sync"
)

type JwtReq struct {
	JwtObj *jwtApi.JwtCfg
	JwtCtx CtxKey
}

type CtxKey struct {
	JwtHeaderKey string
	JwtCtxErrKey string
	JwtCtxMcKey  string
}

// JWTAuthMiddleware
// 基于JWT认证中间件
func JWTAuthMiddleware(req JwtReq) gin.HandlerFunc {
	if req.JwtObj == nil {
		panic("WsHcGnaO req.JwtObj is nil")
	}
	return func(c *gin.Context) {
		jwtStr := c.Request.Header.Get(req.JwtCtx.JwtHeaderKey)
		mc, err := req.JwtObj.ParseToken(jwtStr)
		if err != nil {
			c.Set(req.JwtCtx.JwtCtxErrKey, err)
			return
		}
		//后续的处理函数可以通过c.Get("JwtCtxMcKey")来获取请求的用户信息
		c.Set(req.JwtCtx.JwtCtxMcKey, mc)
	}
}

// GetMcFromCtx
// 从ctx取得 mc
func GetMcFromCtx(ctx *gin.Context, req JwtReq) (mc *jwtApi.MyClaims, err error) {
	t, ok := ctx.Get(req.JwtCtx.JwtCtxMcKey)
	if !ok {
		return nil, errors.New("hLrtXLod JwtCtxMcKey not exists")
	}
	mc = t.(*jwtApi.MyClaims)
	return mc, nil
}

// 从ctx取得JwtCtxErrKey
func GetJwtErrFromCtx(ctx *gin.Context, req JwtReq) error {
	err, ok := ctx.Get(req.JwtCtx.JwtCtxErrKey)
	if !ok {
		return nil
	}
	return err.(error)
}

// GetTokenObjFromCtx
// form ctx 取得 mc 并且将token JsonUnmarshal 到obj中，要求obj为对象地址
func GetTokenObjFromCtx(ctx *gin.Context, req JwtReq, obj any) error {
	mc, err := GetMcFromCtx(ctx, req)
	if err != nil {
		return err
	}
	err = mc.TokenObj(obj)
	if err != nil {
		return err
	}
	return nil
}

// 单例模式
var jwtObj *jwtApi.JwtCfg
var setOnce sync.Once

func InitJwtObj(obj *jwtApi.JwtCfg) {
	if obj == nil {
		panic("GWnRffXb2 InitJwtObj obj nil")
	}
	setOnce.Do(func() {
		jwtObj = obj
	})
}

func GetJwtObj() *jwtApi.JwtCfg {
	if jwtObj != nil {
		return jwtObj
	}
	panic("IsRzQnTh please InitJwtObj first")
}
