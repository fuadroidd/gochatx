package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if tokenstr := ctx.GetHeader("Authorization"); tokenstr != "" {
			if token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			}); err == nil && token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				ctx.Set("username", claims["username"])
				ctx.Next()
			} else {
				ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid token",
				})
				ctx.Abort()

			}
		} else {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			ctx.Abort()
			return
		}

	}
}
