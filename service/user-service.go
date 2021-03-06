package service

import (
	"login/core/repository"
	"login/service/dtos"
	"login/service/mapper"
)

type UserService interface {
	Save(user dtos.UserDTO) (dtos.UserDTO, error)
	Delete(id string) error
	Select(id string) (dtos.UserDTO, error)
	SelectAll() []dtos.UserDTO
}

type userService struct {
	userRepository repository.UserRepository
}


func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) Save(parameter dtos.UserDTO) (dtos.UserDTO, error)  {
	userEntity := mapper.ToUserEntity(parameter)
	user, err := u.userRepository.Save(userEntity)
	userDto := mapper.ToUserDto(user)
	return userDto, err
}

func (u userService) Delete(id string) error {
	err := u.userRepository.Delete(id)
	return err
}

func (u userService) Select(id string) (dtos.UserDTO, error) {
	user, err := u.userRepository.FindById(id)
	userDto := mapper.ToUserDto(user)
	return userDto, err
}

func (u userService) SelectByUsername(username string) (dtos.UserDTO, error) {
	user, err := u.userRepository.FindByUsername(username)
	userDto := mapper.ToUserDto(user)
	return userDto, err
}

func (u userService) SelectAll() []dtos.UserDTO {
	users := u.userRepository.FindAll()
	userDtos := mapper.ToUserDtos(users)
	return userDtos
}