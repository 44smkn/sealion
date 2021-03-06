package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sealion/application/usecase"
	"sealion/domain/model"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var Set = wire.NewSet(NewTaskHandler)

type TaskHandler interface {
	Routes() chi.Router

	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	u usecase.TaskUseCase
}

func NewTaskHandler(u usecase.TaskUseCase) TaskHandler {
	return &taskHandler{u: u}
}

func (h taskHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.Get)
	r.Post("/", h.Create)
	r.Put("/", h.Update)
	r.Delete("/{id}", h.Delete)

	return r 
}

func (h taskHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	tasks, err := h.u.GetTasks(ctx)
	if err != nil {
		logrus.Errorf("failed to get tasks from db.\ndetails: \n%v \n", err)
		respondError(w, http.StatusInternalServerError, "failed to get tasks from db")
	}
	respondWithJson(w, http.StatusOK, tasks)
}

func (h taskHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var task model.Task
	decodeBody(w, r, &task)
	if err := h.u.CreateTask(ctx, task); err != nil {
		logrus.Errorf("failed to create task.\ndetails: \n%v \n", err)
		respondError(w, http.StatusInternalServerError, "failed to create task")
	}
	respondWithJson(w, http.StatusOK, task)
}

func (h taskHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var task model.Task
	decodeBody(w, r, &task)
	if err := h.u.UpdateTask(ctx, task); err != nil {
		logrus.Errorf("failed to update task.\ndetails: \n%v \n", err)
		respondError(w, http.StatusInternalServerError, "failed to update task")
	}
	respondWithJson(w, http.StatusOK, task)
}

func (h taskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 0, 64)
	err := h.u.DeleteTask(ctx, id)
	if err != nil {
		logrus.Errorf("failed to delete task.\ndetails: \n%v \n", err)
		respondError(w, http.StatusInternalServerError, "failed to delete task")
	}
	respondWithJson(w, http.StatusOK, nil)
}

func respondError(w http.ResponseWriter, code int, cause string) {
	respondWithJson(w, code, struct{ cause string }{cause: cause})
}

func respondWithJson(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	body, err := json.Marshal(v)
	if err != nil {
		logrus.Errorf("failed to parse object to json.\ndetails: \n%v \n", err)
	}
	w.Write(body)
}

func decodeBody(w http.ResponseWriter, r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(v); err != nil {
		logrus.Errorf("failed to parse json request body.\ndetails: \n%v \n", err)
		respondError(w, http.StatusBadRequest, "failed to parse json request body")
	}
}
