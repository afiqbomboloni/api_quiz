package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/afiqbomboloni/api_quiz/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.TokenValid(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"errors":  err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func AdminLoginMiddleware() gin.HandlerFunc {
	// cek apakah user yang login adalah admin berdasarkan role, jika bukan maka akan di reject
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		if role != "admin" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"errors":  "You are not authorized to access this resource",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH, HEAD")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeRequest := time.Now()
		method := c.Request.Method
		
		path := c.Request.URL.Path
		c.Next()
		status := c.Writer.Status()
		fmt.Printf("%s - [%s] %s %s %d\n", timeRequest.Format("2006-01-02 15:04:05"), method, path, c.ClientIP(), status)

		
	}

}