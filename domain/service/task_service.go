package service

type TaskService interface {
	SyncJira() error
}

type taskService struct{}

func (s *taskService) SyncJira() error {
	return nil
}
