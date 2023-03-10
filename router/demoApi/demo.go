package demoApi

import (
	"github.com/Godyu97/vege9/jwtApi"
	middleware "github.com/Godyu97/vege9/middleWare"
	"github.com/Godyu97/vege9/vegeTools"
	"github.com/gin-gonic/gin"
)

type DemoPostFnReq struct {
	Id string `json:"Id"`
}
type DemoPostFnResp struct {
	Data string `json:"Data"`
}

func (a Api) DemoPostFn(ctx *gin.Context, req *DemoPostFnReq) (resp *DemoPostFnResp, err error) {
	resp = &DemoPostFnResp{}
	t, _ := ctx.Get(middleware.JwtCtxTokenKey)
	token := t.(*jwtApi.MyClaims)
	resp.Data, _ = vegeTools.JsonMarshalToString(token.AuthData)
	return resp, nil
}
