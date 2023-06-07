package demoApi

import (
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
	//mc, err := midware.GetMcFromCtx(ctx, JwtDefaultReq)
	//if err != nil {
	//	panic(err)
	//}
	////resp token obj json
	//resp.Data, _ = vegeTools.JsonMarshalToString(mc.Token)
	resp.Data = "Hello,vege9!"
	return resp, nil
}
