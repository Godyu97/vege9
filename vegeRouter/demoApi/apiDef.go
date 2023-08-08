package demoApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebError struct {
	code   string
	errMsg string
}

func (err *WebError) Error() string {
	if err != nil {
		return err.errMsg
	}
	return ""
}

func (err *WebError) Code() string {
	if err != nil {
		return err.code
	}
	return ""
}

func GetWebMsgCode(errMsg string) string {
	return errCodeM[errMsg]
}

func GetWebErrorFromErrMsg(errMsg string) *WebError {
	return &WebError{
		code:   GetWebMsgCode(errMsg),
		errMsg: errMsg,
	}
}

var errCodeM = map[string]string{
	//	todo 定义错误码
}

type Api struct {
	UriToFnNameM map[string]string
}

// 所有200 的接口 resp 为ApiResp
type ApiResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty" swaggerignore:"true"`
}

func (a Api) SendOk(c *gin.Context, body any) {
	c.JSON(http.StatusOK, body)
}

func (a Api) SendBad(c *gin.Context, errMsg string) {
	//judge Code
	err := GetWebErrorFromErrMsg(errMsg)
	resp := ApiResp{
		Code: err.Code(),
		Msg:  err.Error(),
		Data: nil,
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}

const CheckAuthErrMsg = "权限不足2333"

func (a Api) CheckAuth(c *gin.Context) error {
	//return nil
	if e, exist := c.Get(JwtDefaultReq.JwtCtx.JwtCtxErrKey); exist {
		msg := fmt.Sprintf("err:%s %s", CheckAuthErrMsg, e)
		err := GetWebErrorFromErrMsg(CheckAuthErrMsg)
		resp := ApiResp{
			Code: err.Code(),
			Msg:  err.Error(),
			Data: msg,
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
