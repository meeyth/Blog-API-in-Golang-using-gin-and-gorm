package middleware

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuthorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, _ := ctx.Cookie("jwt")

		if token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		}); err == nil {

			ctx.Set("claims", token.Claims)
			// ctx.Next()
			return
		}

		ctx.Error(errors.New("someone's trying unauthorized access"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"error": "You are Unauthorized, first Login or Sign Up",
		})
	}
}
