package dtos

import (
	"gorm.io/gorm"
	"time"
)

type UserDTO struct {
	Id uint `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}