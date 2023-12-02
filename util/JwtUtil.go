package util

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Name               string
	jwt.StandardClaims // jwt中标准格式,主要是设置token的过期时间
}

var sign = "student"

func CreateJwt(name string) (string, error) {
	// 当前时间
	nowTime := time.Now()
	// 过期时间
	expireTime := nowTime.Add(300 * time.Second)
	//   签发人
	issuer := "xxx"
	//	 赋值给结构体
	claims := Claims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 转成纳秒
			Issuer:    issuer,
		},
	}
	// 根据签名生成token，NewWithClaims(加密方式,claims) ==》 头部，载荷，签证
	toke, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(sign))
	return toke, err
}

func ParseJwt(token string) (*Claims, error) {
	// ParseWithClaims 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 使用签名解析用户传入的token,获取载荷部分数据
		return []byte(sign), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		//Valid用于校验鉴权声明。解析出载荷部分
		if c, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return c, nil
		}
	}
	return nil, err
}
