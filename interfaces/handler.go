package interfaces

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sealion/application/usecase"
)

type AppHandler interface {
	TaskHandler
}

type TaskHandler interface {
}

type taskHandler struct {
	u usecase.TaskUseCase
}

type errorHandler struct{}

func NewTaskHandler(u usecase.TaskUseCase) TaskHandler {
	return &taskHandler{u}
}

func (t *taskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	switch r.Method {
	case http.MethodGet:
		tasks, err := t.u.GetTasks(ctx)
		if err != nil {
			// something
		}
		respondWithJson(w, http.StatusOK, tasks)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (e *errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func respondWithJson(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var body *bytes.Buffer
	encoder := json.NewEncoder(body)
	err := encoder.Encode(v)
	if err != nil {
		// something
	}
	io.Copy(w, body)
}
