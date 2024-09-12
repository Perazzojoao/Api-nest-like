package repository

import (
	"nest/app/modules/UserModule/entity"
)

type Repository interface {
	Save(user *entity.UserEntity) *entity.UserEntity
	GetUser() *entity.UserEntity
}
