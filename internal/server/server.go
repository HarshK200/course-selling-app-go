package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Addr    string
	Handler http.Handler
}

// return an instance of app takes the port address as input
func NewApp(Address string) *App {
	return &App{
		Addr: Address,
	}
}

// TODO: handlers functions
func (a *App) LoadRoutes() {
	router := chi.NewRouter()

	// middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// routes
	router.Get("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("status ok..."))
	})

	a.Handler = router
}

// starts the application on the port configured during instantitation
func (a *App) Listen() error {
	server := http.Server{
		Addr:    a.Addr,
		Handler: a.Handler,
	}

	log.Printf("App listening on port: %s", a.Addr)
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
