package usermodule

import (
	"go.uber.org/fx"

	"nest/app/modules/UserModule/entity"
	"nest/app/modules/UserModule/repository"
	"nest/nest"
)

var UserModuleFx = fx.Module("userModule",
	fx.Provide(
		NewUserController,
		NewUserService,
		fx.Annotate(
			NewUserModule,
			fx.As(new(nest.Module)),
		),
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(repository.Repository)),
		),
		entity.NewUserEntity,
	),
)

type UserModule struct {
	nest.Module
	modules     []nest.Module
	controllers []nest.Controller
}

func NewUserModule(userController *UserController) *UserModule {
	return &UserModule{
		controllers: []nest.Controller{
			userController,
		},
	}
}

func (um *UserModule) GetControllers() []nest.Controller {
	return um.controllers
}

func (um *UserModule) GetModules() []nest.Module {
	return um.modules
}
