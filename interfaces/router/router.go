package router

import (
	"fmt"
	"net/http"
	"sealion/interfaces/handler"

	"github.com/gorilla/mux"
)

func Run(port int, th handler.TaskHandler) error {

	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", th.Get).Methods("GET")
	r.HandleFunc("/api/tasks", th.Create).Methods("POST")
	r.HandleFunc("/api/tasks", th.Update).Methods("PUT")
	r.HandleFunc("/api/tasks/{id}", th.Delete).Methods("DELETE")
	http.Handle("/", r)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
