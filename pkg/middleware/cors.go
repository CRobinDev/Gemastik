package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Length, Content-Type, Authorization, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
			return
		}

		ctx.Next()
	}
}
