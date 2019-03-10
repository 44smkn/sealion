package interfaces

import (
	"fmt"
	"net/http"
)

func Run(port int) error {
	conn := config.GetDbConn()
	http.Handle("/api/tasks", NewTaskHandler(NewTaskUseCase(NewTaskRepository(conn))))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
