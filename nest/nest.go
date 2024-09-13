package nest

import (
	"net/http"
)

type Handler interface {
	http.Handler
}
