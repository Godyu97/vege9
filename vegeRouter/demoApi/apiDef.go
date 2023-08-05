package demoApi

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	UriToFnNameM map[string]string
}

type ApiResp struct {
	Err  string `json:"Err"`
	Body any    `json:"Body"`
}

func (a Api) SendOk(c *gin.Context, body any) {
	resp := ApiResp{
		Err:  "",
		Body: body,
	}
	c.JSON(http.StatusOK, resp)
}

func (a Api) SendBad(c *gin.Context, errMsg string) {
	resp := ApiResp{
		Err:  errMsg,
		Body: nil,
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}

const CheckAuthErrMsg = "权限不足2333"

func (a Api) CheckAuth(c *gin.Context) error {
	//return nil
	if e, exist := c.Get(JwtDefaultReq.JwtCtx.JwtCtxErrKey); exist {
		err := errors.New(fmt.Sprintf("err:%s %s", CheckAuthErrMsg, e))
		resp := ApiResp{
			Err:  err.Error(),
			Body: nil,
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return err
	}
	return nil
}

func (a Api) UriToFnName(uri string) string {
	fnName, ok := a.UriToFnNameM[uri]
	if ok {
		return fnName
	} else {
		return uri
	}
}
