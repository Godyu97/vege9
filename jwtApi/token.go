package jwtApi

import (
	"github.com/Godyu97/vege9/vegeTools"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Token any
	jwt.RegisteredClaims
}

func (mc *MyClaims) TokenObj(obj any) error {
	return vegeTools.MapToObj(mc.Token.(map[string]any), obj)
}
