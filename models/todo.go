package models

import "time"

type ToDo struct {
	tableName  struct{}  `pg:"todo"`
	Id         int       `json:"id" pg:"id"`
	Title      string    `json:"title" pg:"title"`
	IsComplete bool      `json:"is_complete" pg:"is_complete"`
	UpdateAt   time.Time `json:"update_at" pg:"updated_at"`
	CreatedAt  time.Time `json:"created_at" pg:"created_at"`
}
