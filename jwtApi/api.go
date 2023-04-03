package jwtApi

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
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
		c.ExpiresAt = &jwt.NumericDate{time.Now().Add(j.TokenExpireDuration)}
	}
	if j.Issuer != "" {
		c.Issuer = j.Issuer
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名并获得完成的编码后的字符串token
	return token.SignedString([]byte(j.Key))

}

// ParseToken
// 解析 jwt token
func (j *JwtCfg) ParseToken(tokenStr string) (*MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(j.Key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("sxEAexaD Invalid token")
}
