package router

import (
	"fmt"
	"net/http"
	"sealion/interfaces/handler"
	"sealion/registry"
	"github.com/gorilla/mux"
)

func Run(port int) error {
	h, _ := registry.Store.Get("AppHandler").(handler.TaskHandler)

	r := mux.NewRouter()
	r.Handle("/api/tasks", http.HandlerFunc(h))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
