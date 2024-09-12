package repository

import "nest/app/modules/UserModule/entity"

type UserRepository struct {
	user *entity.UserEntity
}

func NewUserRepository(user *entity.UserEntity) *UserRepository {
	return &UserRepository{
		user: user,
	}
}

func (u *UserRepository) Save(user *entity.UserEntity) *entity.UserEntity {
	u.user = user
	return user
}

func (u *UserRepository) GetUser() *entity.UserEntity {
	return u.user
}
