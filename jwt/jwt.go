package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		[]byte(secretKey),
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(data string, timeout int64) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:  "TimeToken",
		Subject: data,
	}
	if timeout > 0 {
		claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(timeout)).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(token string) (string, error) {
	var err error
	var claims jwt.StandardClaims
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	err = claims.Valid()
	if err != nil {
		return "", err
	} else {
		return claims.Subject, err
	}
}

// MakeTimeoutToken 创建一个会过期的token
func MakeTimeoutToken(data string, secretKey string, timeout int64) (string, error) {
	j := NewJWT(secretKey)
	return j.CreateToken(data, timeout)
}

// ParseToken 解析 token
func ParseToken(secretKey, token string) (string, error) {
	j := NewJWT(secretKey)
	return j.ParseToken(token)
}
