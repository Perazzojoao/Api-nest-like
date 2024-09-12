package nest

import (
	"net/http"

)

type Module interface {
	GetControllers() []Controller
	GetModules() []Module
}

type Controller interface {
	AddRoter() http.Handler
}

type Handler interface {
	http.Handler
}

func AddModule(imports []any, controllers []Controller, providers ...any) {
	for _, c := range imports {
		if m, ok := c.(Module); ok {
			controllers = append(controllers, m.GetControllers()...)
		}
	}
}
