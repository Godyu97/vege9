package authApi

import (
	"log"

	"github.com/Godyu97/vege9/jwtApi"
	middleware "github.com/Godyu97/vege9/middleWare"
	"github.com/Godyu97/vege9/vegeTools"
	"github.com/gin-gonic/gin"
)

type GetUserReq struct {
	Id string `json:"id"`
}
type GetUserResp struct {
	Data string `json:"data"`
}

func (a AuthApi) GetUser(ctx *gin.Context, req *GetUserReq) (resp *GetUserResp, err error) {
	log.Println("lmDMuHes", req)
	resp = &GetUserResp{}
	t, _ := ctx.Get(middleware.JwtCtxTokenKey)
	token := t.(*jwtApi.MyClaims)
	resp.Data, _ = vegeTools.MarshalToString(token.AuthData)
	return resp, nil
}
