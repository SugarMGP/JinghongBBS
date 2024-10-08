package utils

import (
	"BBS/config/config"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

var ErrTokenHandlingFailed = errors.New("token handling failed")

func GenerateToken(userID uint) (string, error) {
	lifespan, err := strconv.Atoi(config.Config.GetString("jwt.lifespan"))
	if err != nil {
		return "", err
	}

	claims := UserClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(lifespan) * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                          // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.GetString("jwt.secret")))
	return tokenString, err
}

// 从 jwt 中解析出 user_id
func ExtractToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(config.Config.GetString("jwt.secret")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*UserClaims)
	// 如果 jwt 有效，返回 user_id
	if ok && token.Valid {
		return claims.UserID, nil
	}
	return 0, ErrTokenHandlingFailed
}
