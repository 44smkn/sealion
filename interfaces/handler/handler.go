package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sealion/application/usecase"
)

type AppHandler interface {
	TaskHandler
}

type TaskHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	u usecase.TaskUseCase
}

type errorHandler struct {
	cause string
	code  int
}

func NewTaskHandler(u usecase.TaskUseCase) TaskHandler {
	return &taskHandler{u}
}

func (t *taskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	switch r.Method {
	case http.MethodGet:
		tasks, err := t.u.GetTasks(ctx)
		if err != nil {
			log.Println(err)
			eh := &errorHandler{
				cause: "failed to get tasks from db",
				code:  http.StatusInternalServerError,
			}
			eh.ServeHTTP(w, r)
		}
		respondWithJson(w, http.StatusOK, tasks)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (e *errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, e.code, struct{ cause string }{cause: e.cause})
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