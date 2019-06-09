package service_test

import (
	"context"
	"sealion/domain/model"
	"sealion/domain/service"
	"sealion/infrastructure/client"
	"testing"
)

func TestSyncJira(t *testing.T) {

	c, err := client.NewJira()
	if err != nil {
		t.Error(err)
	}
	s := service.NewTaskService(c)
	ctx := context.Background()
	existedTasks := []*model.Task{}
	_, err = s.SyncJira(ctx, existedTasks)
	if err != nil {
		t.Error(err)
	}
}
