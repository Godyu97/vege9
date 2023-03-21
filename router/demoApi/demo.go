package demoApi

import (
	"github.com/Godyu97/vege9/middleware"
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
	mc, err := middleware.GetMcFromCtx(ctx)
	if err != nil {
		panic(err)
	}
	//resp token obj json
	resp.Data, _ = vegeTools.JsonMarshalToString(mc.AuthData)
	return resp, nil
}
