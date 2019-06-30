package router

import (
	"fmt"
	"net/http"
	"sealion/interfaces/handler"

	"github.com/go-chi/chi"
)

func Run(port int, th handler.TaskHandler) error {

	r := chi.NewRouter()
	r.Route("/api/tasks", func(r chi.Router) {
		r.Get("/", th.Get)
		r.Post("/", th.Create)
		r.Put("/", th.Update)
		r.Delete("/{id}", th.Delete)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
