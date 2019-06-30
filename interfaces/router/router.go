package router

import (
	"github.com/go-chi/chi/middleware"
	"fmt"
	"net/http"
	"sealion/interfaces/handler"

	"github.com/go-chi/chi"
)

func Run(port int, th handler.TaskHandler) error {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api/tasks", th.Routes())
	
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
