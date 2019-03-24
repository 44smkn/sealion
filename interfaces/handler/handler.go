package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sealion/application/usecase"
	"sealion/domain/model"
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
	case http.MethodPost:
		var task model.Task
		decodeBody(w, r, &task)
		if err := t.u.CreateTask(ctx, task); err != nil {
			log.Println(err)
			eh := &errorHandler{
				cause: "failed to create task",
				code:  http.StatusInternalServerError,
			}
			eh.ServeHTTP(w, r)
		}
		respondWithJson(w, http.StatusOK, task)
	case http.MethodPut:
		var task model.Task
		decodeBody(w, r, &task)

		if err := t.u.UpdateTask(ctx, task); err != nil {
			log.Println(err)
			eh := &errorHandler{
				cause: "failed to update task",
				code:  http.StatusInternalServerError,
			}
			eh.ServeHTTP(w, r)
		}
		respondWithJson(w, http.StatusOK, task)
	default:
		respondWithJson(w, http.StatusNotFound, fmt.Sprintf("{\"cause\": \"method %v is not found\"}", r.Method))
	}
}

func (e *errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, e.code, struct{ cause string }{cause: e.cause})
}

func respondWithJson(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	body, err := json.Marshal(v)
	if err != nil {
		// something
	}
	w.Write(body)
}

func decodeBody(w http.ResponseWriter, r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(v); err != nil {
		log.Println(err)
		eh := &errorHandler{
			cause: "failed to parse json request body",
			code:  http.StatusBadRequest,
		}
		eh.ServeHTTP(w, r)
	}
}
