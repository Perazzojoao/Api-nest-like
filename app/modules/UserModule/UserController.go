package usermodule

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"

	"nest/app/modules/UserModule/entity"
)

type UserController struct {
	userService *UserService
	Router      *chi.Mux
}

var tools = toolbox.New()

func NewUserController(userService *UserService) *UserController {
	uc := &UserController{
		userService: userService,
	}
	uc.AddRoter()
	return uc
}

func (u *UserController) AddRoter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Get("/", u.GetUser)
		r.Post("/", u.AddUser)
	})

	return r
}

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	tools.WriteJSON(w, http.StatusOK, u.userService.GetUser())
}

func (u *UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	var payload *entity.UserEntity
	err := tools.ReadJSON(w, r, payload)
	if err != nil {
		tools.ErrorJSON(w, errors.New("error reading JSON"), http.StatusBadRequest)
		return
	}

	tools.WriteJSON(w, http.StatusOK, u.userService.SaveUser(payload))
}
