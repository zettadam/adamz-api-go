package models

import (
	"database/sql"
	"time"
)

type Task struct {
	Id          int64          `json:"id"`
	TaskId      sql.NullInt64  `json:"taskId"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}
