package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	"nest/app"
	"nest/nest"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			newHTTPServer,
			fx.Annotate(
				newChiHandler,
				fx.As(new(nest.Handler)),
			),
		),
		app.AppModuleFx,
		fx.Invoke(func(*http.Server) {}),
	)
}

func newHTTPServer(lc fx.Lifecycle, r nest.Handler) *http.Server {
	srv := &http.Server{Addr: ":8000", Handler: r}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func newChiHandler(appModule *app.AppModule) *chi.Mux {
	r := chi.NewRouter()
	for _, c := range appModule.GetControllers() {
		r.Mount("/", c.AddRoter())
	}
	return r
}
