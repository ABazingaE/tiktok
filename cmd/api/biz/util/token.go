package util

import (
	"github.com/form3tech-oss/jwt-go"
	"strconv"
	"tiktok/pkg/consts"
)

// 解析token中的user_id
func GetUserIdFromToken(token string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	var requestUserId interface{}
	for key, val := range claims {
		if key == "user_id" {
			requestUserId = val
			break
		}
	}
	return strconv.FormatFloat(requestUserId.(float64), 'f', -1, 64), nil
}
