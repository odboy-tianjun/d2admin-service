package jwt

import (
	"d2-admin-service/src/infra/config"
	"github.com/fatih/color"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// CustomClaims 定义自定义claims结构体
type CustomClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

// GenToken 生成token
func GenToken(username string) string {
	secret := config.Config.Jwt.Secret
	// 私钥（用于HS256签名时用作secret，对于RS256等非对称算法则是私钥）
	key := []byte(secret)

	// 生成claims
	claims := &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "odboy.cn",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 设置过期时间
			Subject:   username,                                           // 用户ID或其他唯一标识符
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Username: username,
	}

	// 创建一个新的token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥进行签名并获取完整的编码后的token
	signedToken, err := token.SignedString(key)
	if err != nil {
		panic("Generated JWT Error: " + err.Error())
	}
	color.Green("%s Generated JWT: %s\n", username, signedToken)
	return signedToken
}

// ParseToken 解析token
func ParseToken(signedToken string) (int, *CustomClaims) {
	// 解析JWT
	parser := jwt.Parser{}
	// 需要设置Valid方法以验证claims中的标准字段，例如ExpiresAt
	var parsedClaims *CustomClaims // 将parsedClaims声明为指针类型
	_, _, err := parser.ParseUnverified(signedToken, parsedClaims)
	if err != nil {
		//panic("无效Token" + err.Error())
		color.Red("无效token, %v\n", err)
		return -1, nil
	}
	// 如果需要验证签名，请使用正确的秘钥和方法
	secret := config.Config.Jwt.Secret
	key := []byte(secret)
	verifiedToken, err := parser.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		color.Red("无效token, %v\n", err)
		return -1, nil
	}
	parsedClaims, ok := verifiedToken.Claims.(*CustomClaims)
	if !ok || !verifiedToken.Valid {
		//panic("Invalid token")
		color.Red("token未通过校验, %v\n", err)
		return -1, nil
	}
	return 200, parsedClaims
}
