package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	Category    string    `json:"category"`
	Name        string    `json:"name"`
	DoToday     bool      `json:"doToday"`
	Deadline    time.Time `json:"deadline"`
	Description string    `json:"description"`
	TicketId    string    `json:"ticketId,omitempty"`
	Archive     bool      `json:"archive"`
}
