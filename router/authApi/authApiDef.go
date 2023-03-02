package authApi

import (
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
