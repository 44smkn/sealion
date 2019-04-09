package service_test

import (
	"context"
	"log"
	"os"
	"sealion/domain/model"
	"sealion/domain/service"
	"sealion/infrastructure/client"
	"testing"
)

func TestSyncJira(t *testing.T) {

	baseUrl := os.Getenv("JIRA_BASE_URL")
	username := os.Getenv("JIRA_USERNAME")
	password := os.Getenv("JIRA_PASSWORD")
	c, err := client.NewJira(baseUrl, username, password)
	if err != nil {
		log.Println(err)
	}
	s := service.NewTaskService(*c)
	ctx := context.Background()
	existedTasks := []*model.Task{}
	_, err = s.SyncJira(ctx, existedTasks)
	if err != nil {
		t.Error(err)
	}
}
