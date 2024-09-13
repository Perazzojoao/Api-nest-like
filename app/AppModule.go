package app

import (
	"go.uber.org/fx"

	usermodule "nest/app/modules/UserModule"
	"nest/nest"

)

var AppModuleFx = fx.Module("appModule",
	fx.Provide(
		fx.Annotate(
			NewAppModule,
			fx.ParamTags(`group:"appModules"`),
		),
	),
	usermodule.UserModuleFx,
)

type AppModule struct {
	controllers []nest.Controller
}

func NewAppModule(module []nest.IModule) *AppModule {
	appModule := &AppModule{}

	for _, module := range module {
		appModule.controllers = append(appModule.controllers, module.GetControllers()...)
	}
	return appModule
}

func (app *AppModule) GetControllers() []nest.Controller {
	return app.controllers
}
