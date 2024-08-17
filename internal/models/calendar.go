package models

import (
	"database/sql"
	"time"
)

type Event struct {
	Id          int64          `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"slug"`
	StartTime   sql.NullTime   `json:"start_time"`
	EndTime     sql.NullTime   `json:"end_time"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}
