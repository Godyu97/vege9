package authApi

import (
	"log"

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
	resp.Data = "i am vege9"
	return resp, nil
}
