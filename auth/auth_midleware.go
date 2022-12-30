package auth

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ramdanariadi/dot-test/helpers"
	"strings"
)

//go:embed JWTSECRET
var token_secret []byte

func SecureRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Authorization")

		if strings.HasPrefix(apiKey, "Bearer ") {
			token := strings.SplitAfter(apiKey, "Bearer ")[1]
			decodedJwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
					return token_secret, nil
				}
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			})
			helpers.PanicIfError(err)
			if claims, ok := decodedJwt.Claims.(jwt.MapClaims); ok && decodedJwt.Valid {

				if len(claims) > 0 {
					if claims["role"] == "user" {
						c.Next()
					}
				}
				panic("NOT_AUTHENTICATED")
			}
		}
		panic("NOT_AUTHENTICATED")
	}
}
