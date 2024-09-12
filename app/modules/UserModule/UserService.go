package usermodule

import (
	"nest/app/modules/UserModule/entity"
	"nest/app/modules/UserModule/repository"
)

type UserService struct {
	userRepository repository.Repository
}

func NewUserService(userRepository repository.Repository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) SaveUser(user *entity.UserEntity) *entity.UserEntity {
	return u.userRepository.Save(user)
}

func (u *UserService) GetUser() *entity.UserEntity {
	return u.userRepository.GetUser()
}
