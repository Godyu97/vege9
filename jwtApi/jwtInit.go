package jwtApi

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	AuthData interface{}
	jwt.RegisteredClaims
}

type JwtCfg struct {
	Key                 string
	TokenExpireDuration time.Duration
	Issuer              string
}

type OptionFunc func(cfg *JwtCfg)

// InitJwt
// 初始化jwt管理器
func InitJwt(secret string, opts ...OptionFunc) *JwtCfg {
	o := &JwtCfg{Key: secret}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// WithExp
// 过期时间
func WithExp(exp time.Duration) OptionFunc {
	return func(o *JwtCfg) {
		o.TokenExpireDuration = exp
	}
}

// WithIssuer
// 添加issuer 项目名称
func WithIssuer(iss string) OptionFunc {
	return func(o *JwtCfg) {
		o.Issuer = iss
	}
}
