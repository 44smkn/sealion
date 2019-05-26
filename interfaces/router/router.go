package router

import (
	"fmt"
	"net/http"
	"sealion/interfaces/handler"
	"sealion/registry"
	"github.com/gorilla/mux"
)

func Run(port int) error {

	h, _ := registry.Store.Get("Handler").(handler.Handler)
	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", h.GetTasks).Methods("GET")
	r.HandleFunc("/api/tasks", h.CreateTasks).Methods("POST")
	r.HandleFunc("/api/tasks", h.UpdateTasks).Methods("PUT")
	r.HandleFunc("/api/tasks/{id}", h.DeleteTasks).Methods("DELETE")

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
