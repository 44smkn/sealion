package router

import (
	"fmt"
	"net/http"
	"sealion/registry"
)

func Run(port int) error {
	r := registry.NewRegistry()
	http.Handle("/api/tasks", r.NewAppHandler())
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
