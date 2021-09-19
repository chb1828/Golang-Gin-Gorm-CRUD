package mapper

import (
	"login/core/entity"
	"login/service/dtos"
)

func ToUserEntity(dto dtos.UserDTO) entity.User {
	user := entity.User{
		Username: dto.Username,
		Phone: dto.Phone,
	}
	user.SetNewPassword(dto.Password)
	return user
}

func ToUserDto(user entity.User) dtos.UserDTO {
	userDto := dtos.UserDTO{
		Id: user.ID,
		Username: user.Username,
		Password: string(user.Password),
		Phone: user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	return userDto
}

func ToUserDtos(users []entity.User) []dtos.UserDTO {
	var result []dtos.UserDTO
	for _, value := range users {
		userDto := dtos.UserDTO{
			Id: value.ID,
			Username: value.Username,
			Password: string(value.Password),
			Phone: value.Phone,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			DeletedAt: value.DeletedAt,
		}
		result = append(result,userDto)
	}
	return result
}
