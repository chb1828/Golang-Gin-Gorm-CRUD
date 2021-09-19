package service

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"login/core/entity"
	"login/service/dtos"
	"os"
	"time"
)

var user = entity.User{
	Username: "admin",
	Password: []byte("admin"),
	Phone:    "49123454322",
}

type LoginService interface {
	CreateToken(dto dtos.LoginDTO) (string,error)
	Check(dto dtos.LoginDTO) bool
}

type loginService struct {
	token string
}

func NewLoginService() LoginService {
	return &loginService{}
}
func (service *loginService) CreateToken (dto dtos.LoginDTO) (string,error) {
	var err error

	// 일시적으로 아래와 같이 환경 변수를 설정한다
	os.Setenv("ACCESS_SECRET","test")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = dto.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}
	return token, nil
}

func (service *loginService) Check(dto dtos.LoginDTO) bool {
	// 실제로 DB에 요청한 계정이 유효한지 검증한다
	if user.Username != dto.Username || checkPasswordHash(user.Password,dto.Password) {
		return false
	}
	return true
}

func checkPasswordHash(password []byte, hash string) bool {
	err := bcrypt.CompareHashAndPassword(password,[]byte(hash))
	return err == nil
}
