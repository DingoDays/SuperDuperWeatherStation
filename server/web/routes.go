package web

import (
	"github.com/go-chi/chi/v5"
	"server/core"
)

type Routes struct {
	App      *core.Application
	Handlers *Handler
}

func NewRouter(app *core.Application) *Routes {
	routes := &Routes{}

	routes.App = app
	routes.Handlers = NewHandler(app)

	return routes
}

func (r *Routes) HandleRequests() *chi.Mux {
	router := chi.NewRouter()

	//GET
	router.Get("/ping", r.Handlers.Ping)
	router.Get("/get", r.Handlers.Get)

	// POST
	router.Post("/add", r.Handlers.Add)

	return router
}
