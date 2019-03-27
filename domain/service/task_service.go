package service

import (
	"context"
	"log"
	"sealion/infrastructure/client"
)

type TaskService interface {
	SyncJira(ctx context.Context) error
}

type taskService struct {
	client.JiraClient
}

func NewTaskService(c client.JiraClient) TaskService {
	return &taskService{c}
}

func (s *taskService) SyncJira(ctx context.Context) error {
	issues, err := s.GetMyIssue(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(issues)
	return nil
}
