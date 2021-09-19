package main

import (
	"github.com/gin-gonic/gin"
	"login/config"
	"login/core/database"
	"login/route"
)

func main() {
	//환경변수 초기화
	config.InitEnv()
	//데이터베이스 초기화
	database.InitDB()
	//서버 초기화
	server := gin.New()

	//라우터 초기화
	route.InitRoutes(server)

	server.Run(":8080")
}
