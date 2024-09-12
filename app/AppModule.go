package app

import (
	"go.uber.org/fx"

	usermodule "nest/app/modules/UserModule"
	"nest/nest"
)

var AppModuleFx = fx.Module("appModule",
	fx.Provide(NewAppModule),
	usermodule.UserModuleFx,
)

type AppModule struct {
	modules     []nest.Module
	controllers []nest.Controller
}

func NewAppModule(module nest.Module) *AppModule {
	appModule := &AppModule{}

	appModule.modules = append(appModule.modules, module.GetModules()...)
	appModule.controllers = append(appModule.controllers, module.GetControllers()...)
	return appModule
}

func (app *AppModule) GetControllers() []nest.Controller {
	return app.controllers
}

func (app *AppModule) GetModules() []nest.Module {
	return app.modules
}
