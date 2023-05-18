package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
)

type Harvester struct {
	ctx    context.Context
	router *chi.Mux
	logger *slog.Logger
	cfg    Config
}

type Config struct {
	Env  string
	Port int
}

type Option func(*Harvester)

func WithConfig(cfg Config) Option {
	return func(app *Harvester) {
		app.cfg = cfg
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(app *Harvester) {
		app.logger = logger
	}
}

func RegisterUsecase[T any](method, version, path string, handler Handler[T]) Option {
	return func(app *Harvester) {
		app.router.MethodFunc(method, fmt.Sprintf("/%s/%s", version, path), createHandler(app, handler))
	}
}

func New(opts ...Option) *Harvester {
	r := chi.NewRouter()

	app := &Harvester{
		router: r,
		ctx:    context.Background(),
	}

	for _, use := range opts {
		use(app)
	}

	return app
}

func (app *Harvester) Run() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.Port),
		Handler:      app.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     nil,
	}
	app.logger.Info(
		"harvester API started",
		slog.String("env", app.cfg.Env),
		slog.Int("port", app.cfg.Port),
	)

	return srv.ListenAndServe()
}
