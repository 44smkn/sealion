package model

import "time"

type Task struct {
	Id       int       `json:"id"`
	Category string    `json:"category"`
	Name     string    `json:"name"`
	DoToday  bool      `json:"doToday"`
	Deadline time.Time `json:"deadline"`
}
