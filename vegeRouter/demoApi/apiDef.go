package demoApi

import (
	"errors"
	"fmt"
	"github.com/Godyu97/vege9/midware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ApiObj Api

type ApiResp struct {
	Err  string `json:"Err"`
	Body any    `json:"Body"`
}

type Api struct{}

func (a Api) SendOk(c *gin.Context, body any) {
	resp := ApiResp{
		Err:  "",
		Body: body,
	}
	c.JSON(http.StatusOK, resp)
}

func (a Api) SendBad(c *gin.Context, errMsg string, body any) {
	resp := ApiResp{
		Err:  errMsg,
		Body: body,
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}

const CheckAuthErrMsg = "权限不足2333"

func (a Api) CheckAuth(c *gin.Context, body any) error {
	if e, exist := c.Get(midware.JwtCtxErrKey); exist {
		err := errors.New(fmt.Sprintf("err:%s %s", CheckAuthErrMsg, e))
		resp := ApiResp{
			Err:  err.Error(),
			Body: body,
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return err
	}
	return nil
}
