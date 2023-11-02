package jwtApi

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// SignedTokenStr
// 签发jwt token
func (j *JwtCfg) SignedTokenStr(data any) (string, error) {
	c := MyClaims{
		data,
		jwt.RegisteredClaims{},
	}
	if j.TokenExpireDuration != 0 {
		c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(j.TokenExpireDuration))
	}
	if j.Issuer != "" {
		c.Issuer = j.Issuer
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名并获得完成的编码后的字符串token
	return token.SignedString([]byte(j.key))

}

// SignedTokenStrWithID
// 签发jwt token with ID
func (j *JwtCfg) SignedTokenStrWithID(data any, id string) (string, error) {
	c := MyClaims{
		data,
		jwt.RegisteredClaims{},
	}
	if j.TokenExpireDuration != 0 {
		c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(j.TokenExpireDuration))
	}
	if j.Issuer != "" {
		c.Issuer = j.Issuer
	}
	if id != "" {
		c.ID = id
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名并获得完成的编码后的字符串token
	return token.SignedString([]byte(j.key))

}

const (
	PrefixBearer = "Bearer"
)

// ParseToken
// 解析 jwt token
func (j *JwtCfg) ParseToken(tokenStr string) (*MyClaims, error) {
	//去掉一些奇怪的东西
	i := strings.IndexFunc(tokenStr, func(r rune) bool {
		return r == 32
	})
	tokenStr = tokenStr[i:]
	//解析token
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i any, err error) {
		return []byte(j.key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("sxEAexaD Invalid token")
}
