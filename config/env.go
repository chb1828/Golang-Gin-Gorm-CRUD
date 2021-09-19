package config

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnv()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("환경변수를 로드하는 과정에서 에러가 발생했습니다")
	}
}
