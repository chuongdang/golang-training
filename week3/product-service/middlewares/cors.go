package middlewares

import "github.com/gin-gonic/gin"

/*
CORSMiddleware to handle CORS permissions
*/
func CORSMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Access-Token")
	if ctx.Request.Method == "OPTIONS" || ctx.Request.Method == "HEAD" {
		ctx.AbortWithStatus(200)
		return
	}
	ctx.Next()
}
