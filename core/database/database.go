package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"login/core/entity"
	"os"
)

var DB *gorm.DB

func InitDB() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	db,_ := connectToDB(user,password,dbName,port)
	if db != nil {
		autoMigrate()
	}
}

func connectToDB(dbUser string, dbPassword string, dbName string, port string) (*gorm.DB, error) {
	var dsn = fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		dbUser, dbPassword, dbName, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db !=nil {
		DB = db
	}
	return db, err
}

func GetDB() *gorm.DB {
	return DB
}

func autoMigrate() {
	DB.AutoMigrate(&entity.User{})
}