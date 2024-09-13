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
			fx.ResultTags(`group:"appModules"`),
			fx.As(new(nest.IModule)),
		),
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(repository.Repository)),
		),
		entity.NewUserEntity,
	),
)

type UserModule struct {
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
