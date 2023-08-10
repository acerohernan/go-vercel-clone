package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/acerohernan/go-vercel-clone/api-service/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header["Authorization"]

		if len(header) != 1 {
			ctx.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		bearer := header[0]

		token, prefixFound := strings.CutPrefix(bearer, "Bearer ")

		if !prefixFound {
			ctx.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)

		if !ok || !parsedToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		ctx.Set("user", config.JWTUser{
			Id:    int(claims["sub"].(float64)),
			Email: claims["email"].(string),
		})
		ctx.Next()
	}
}
