package util

import (
	"time"

	er "glossika/internal/domain/er"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GetJwtSecretKey() []byte {
	return []byte(viper.GetString("jwt-secret-key"))
}

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(GetJwtSecretKey())
	if err != nil {
		return "", er.W(err)
	}
	return "Bearer " + tokenString, nil
}
