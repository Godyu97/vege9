package jwtApi

import (
	"testing"
	"time"
)

type TestUser struct {
	Id     string
	Name   string
	Phone  string
	Family FamilyDef
}
type FamilyDef struct {
	Dad string
	Mom string
}

func TestJwt(t *testing.T) {
	key := "RPJbreTM"
	jwtObj := InitJwt(key,
		WithIssuer("TestJwt"),
		WithExp(time.Hour),
	)
	family := FamilyDef{
		Dad: "Fu",
		Mom: "Hui",
	}
	u := TestUser{
		Id:     "0001",
		Name:   "hongyu",
		Phone:  "133344445555",
		Family: family,
	}
	token, err := jwtObj.SignedTokenStr(u)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)
	mc, err := jwtObj.ParseToken(token)
	if err != nil {
		t.Error(err)
		return
	}
	a := &TestUser{}
	err = mc.TokenObj(a)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(a)
	t.Log(mc)
}
