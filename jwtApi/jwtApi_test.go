package jwtApi

import (
	"github.com/Godyu97/vege9/vegeTools"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	key := "godyu"
	jwtObj := InitJwt(key,
		WithIssuer("hongyu"),
		WithExp(time.Hour),
	)
	type user struct {
		Id    string
		Name  string
		Phone string
	}
	u := user{
		Id:    "0001",
		Name:  "hongyu",
		Phone: "133",
	}
	token, err := jwtObj.SignedTokenStr(u)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)
	parseToken, err := jwtObj.ParseToken(token)
	if err != nil {
		t.Error(err)
		return
	}
	a := &user{}
	err = vegeTools.MapToObj(parseToken.AuthData.(map[string]interface{}), a)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(a)
	t.Log(parseToken)
}
