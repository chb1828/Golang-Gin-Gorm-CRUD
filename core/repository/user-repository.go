package repository

import (
	"gorm.io/gorm"
	"log"
	"login/core/entity"
)

type userRepository struct {
 	DB *gorm.DB
}

type UserRepository interface {
	Save(user entity.User)(entity.User, error)
	Delete(id string) error
	FindById(id string) entity.User
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB:db,
	}
}

func (u userRepository) Save(user entity.User) (entity.User, error) {
	log.Println("[UserRepository]...Save 호출")
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) Delete(id string) error {
	log.Println("[UserRepository]...Delete 호출")
	err := u.DB.Delete(&entity.User{},id).Error
	return err
}

func (u userRepository) FindById(id string) entity.User {
	var user entity.User
	log.Println("[UserRepository]...FindById 호출")
	u.DB.First(&user,id)
	return user
}