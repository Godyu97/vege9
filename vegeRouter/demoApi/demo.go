package demoApi

import (
	"github.com/Godyu97/vege9/midware"
	"github.com/Godyu97/vege9/vege"
	"github.com/gin-gonic/gin"
)

type DemoFnReq struct {
	Id string `json:"id" form:"id"`
}
type DemoFnResp struct {
	Id   string `json:"id" form:"id"`
	Data string `json:"Data"`
}

func (a Auth) DemoPostFn(ctx *gin.Context, req *DemoFnReq) (resp *DemoFnResp, err error) {
	resp = &DemoFnResp{}
	mc, err := midware.GetMcFromCtx(ctx, JwtDefaultReq)
	if err != nil {
		panic(err)
	}
	//resp token obj json
	resp.Data, _ = vege.JsonMarshalToString(mc.Token)
	resp.Id = req.Id
	resp.Data = "Hello,vege9!Post"
	return resp, nil
}

func (a Auth) DemoGetFn(ctx *gin.Context, req *DemoFnReq) (resp *DemoFnResp, err error) {
	resp = &DemoFnResp{}
	resp.Id = req.Id
	resp.Data = "Hello,vege9!Get"
	return resp, nil
}

func (a Customer) VegeApi(ctx *gin.Context, req *struct{}) (ApiResp, error) {
	var ptr *string
	//panic
	return ApiResp{
		Code: "",
		Msg:  "",
		Data: *ptr,
	}, nil
}
