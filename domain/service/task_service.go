package service

import (
	"context"
	"log"
	"sealion/domain/model"
	"sealion/infrastructure/client"

	"github.com/google/wire"
)

var Set = wire.NewSet(NewTaskService)

type TaskService interface {
	SyncJira(ctx context.Context, existedTasks []*model.Task) ([]*model.Task, error)
}

type taskService struct {
	*client.JiraClient
}

func NewTaskService(c *client.JiraClient) TaskService {
	return &taskService{c}
}

func (s *taskService) SyncJira(ctx context.Context, existedTasks []*model.Task) ([]*model.Task, error) {
	issues, err := s.GetMyIssue(ctx)
	if err != nil {
		log.Println(err)
	}
	var tasks []*model.Task
	for _, i := range issues {
	LOOP:
		for _, e := range existedTasks {
			if i.ID == e.TicketId {
				break LOOP
			}
		}
		t := &model.Task{
			Category:    "TICKET",
			Name:        i.Fields.Issuetype.Name,
			Description: i.Fields.Summary,
			// Deadline:    i.Fields.DueDate,
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
