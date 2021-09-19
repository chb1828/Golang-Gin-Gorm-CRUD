package middleware

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Logger() gin.HandlerFunc {

	LOGFilEName := "logs"
	// 로그를 기록한다
	if _, err := os.Stat(LOGFilEName); os.IsNotExist(err) {
		err := os.Mkdir(LOGFilEName,0755)
		if err != nil {
			log.Fatal(err)
		}
		logFile, _ := os.Create("log/server.log")
		gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	}

	return gin.Logger()
}