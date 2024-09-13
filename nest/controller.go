package nest

import (
	"github.com/go-chi/chi/v5"
)

type Controller interface {
	AddRoter() *chi.Mux
}
