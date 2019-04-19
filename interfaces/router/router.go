package router

import (
	"fmt"
	"net/http"
	"sealion/interfaces/handler"
	"sealion/registry"
)

func Run(port int) error {
	h, _ := registry.Store.Get("AppHandler").(handler.TaskHandler)
	http.Handle("/api/tasks", h)
	http.Handle("/api/tasks/", h)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
