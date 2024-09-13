package nest

type IModule interface {
	GetControllers() []Controller
}

type Module struct {
	controllers []Controller
}

func (m *Module) GetControllers() []Controller {
	return m.controllers
}
