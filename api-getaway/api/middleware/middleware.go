package middleware

import (
	"api_getaway/api/token"
	"api_getaway/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		fmt.Print("auth: ", auth)
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": "authorization header is required"})
			return
		}

		fmt.Println(config.Load().SIGNING_KEY)
		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid token: %s", err)})
			return
		}

		claims, err := token.ExtractClaims(auth)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid token claims: %s", err)})
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
