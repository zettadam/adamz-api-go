package types

import (
	"time"
)

type Task struct {
	Id          int64      `db:"id"`
	TaskId      *int64     `db:"task_id"`
	Title       string     `db:"title"`
	Description *string    `db:"description"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type TaskRequest struct {
	Id          *int64     `json:"id"`
	TaskId      *int64     `json:"task_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
