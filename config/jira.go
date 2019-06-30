package config

import (
	"sealion/infrastructure/client"
	"github.com/kelseyhightower/envconfig"
)

type JiraConfig struct {
	SyncIssue bool `default:"false split_words:"true"`
	Username string 
	Password string
	BaseUrl string `split_words:"true"`
}


func NewJiraClient() (*client.JiraClient, error){
	var j JiraConfig
	err := envconfig.Process("jira", &j)
	if err != nil {
		return nil, err
	}

	if !j.SyncIssue {
		return &client.JiraClient{}, nil
	}

	client, err := client.NewJira(j.Username, j.Password, j.BaseUrl)
	if err != nil {
		return nil, err
	}

	return client, nil
}