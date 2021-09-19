package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255"`
	Password []byte
	Phone string `gorm:"type:varchar(100)"`
}

func (user *User) SetNewPassword(passwordString string) {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordString),bcrypt.DefaultCost)
	user.Password = bcryptPassword
}