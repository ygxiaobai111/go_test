package dto

import "OceanLearn/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.Go_gin_User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
