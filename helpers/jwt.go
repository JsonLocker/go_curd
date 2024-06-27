package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
}

// 创建JWT // https://pkg.go.dev/github.com/golang-jwt/jwt/v5#NewWithClaims
func JwtCreate(keyword string, expireHour int) string {
	claims := MyCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHour) * time.Hour)),
			Issuer:    keyword,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := []byte(os.Getenv("MMAC_SHA_KEY"))
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return err.Error()
	}
	return tokenString
}

// 解析JWT  // https://pkg.go.dev/github.com/golang-jwt/jwt/v5#ParseWithClaims
func JwtParse(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("MMAC_SHA_KEY")), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return "", err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims.RegisteredClaims.Issuer, nil
	} else {
		return "", errors.New("Invalid token")
	}
}
