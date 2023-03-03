package authApi

import (
	"errors"
	middleware "github.com/Godyu97/vege9/middleWare"
	"github.com/gin-gonic/gin"
	"net/http"
)

var AuthApiObj AuthApi

type AuthResp struct {
	Err  string      `json:"err"`
	Body interface{} `json:"body"`
}

type AuthApi struct{}

func (a AuthApi) SendOk(c *gin.Context, body any) {
	resp := AuthResp{
		Err:  "",
		Body: body,
	}
	c.JSON(http.StatusOK, resp)
}

func (a AuthApi) SendBad(c *gin.Context, errMsg string, body any) {
	resp := AuthResp{
		Err:  errMsg,
		Body: body,
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}

const CheckAuthErrMsg = "权限不足2333"

func (a AuthApi) CheckAuth(c *gin.Context) error {
	if _, exist := c.Get(middleware.JwtCtxErrKey); exist {
		return errors.New(CheckAuthErrMsg)
	}
	return nil
}
