package model

type Tasks []Task

type Task struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	DoToday  bool   `json:"doToday"`
	Deadline Date   `json:"deadline"`
}
