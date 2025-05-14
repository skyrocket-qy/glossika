package middleware

import (
	"time"

	"recsvc/internal/domain/er"
	"recsvc/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			er.Bind(c, er.NewAppErr(er.Unauthorized))
			return
		}

		tokenString := authHeader[7:]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return util.GetJwtSecretKey(), nil
		})
		if err != nil || !token.Valid {
			er.Bind(c, er.NewAppErr(er.Unauthorized))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			er.Bind(c, er.NewAppErr(er.Unauthorized))
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			er.Bind(c, er.NewAppErr(er.Unauthorized))
			return
		}

		c.Set("userID", claims["userID"])

		c.Next()
	}
}
