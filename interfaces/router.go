package interfaces

import (
	"fmt"
	"net/http"
)

func Run(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
